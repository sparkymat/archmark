import React, { useCallback, useMemo } from 'react';
import { useParams } from 'react-router-dom';
import {
  ActionIcon,
  Box,
  Container,
  Flex,
  LoadingOverlay,
  Menu,
  Paper,
  Text,
  Title,
  rem,
} from '@mantine/core';
import { IconSettings, IconTrash } from '@tabler/icons-react';
import { modals } from '@mantine/modals';
import { EditableImage } from '../EditableImage';
import { EditableTextField } from '../EditableTextField';
import { Bookmark } from '../../models/Bookmark';
import {
  useShowQuery,
  useUploadImageMutation,
  useUpdateTitleMutation,
  useUpdateTranscriptMutation,
  useDestroyMutation,
} from '../../slices/Bookmark';
import {
  useCreateMutation as useCreateNoteMutation,
  useRecentQuery as useRecentNotesQuery,
} from '../../slices/AnalysisNote';
import { AnalysisPanel } from '../AnalysisPanel';
import { AnalysisNote } from '../../models/AnalysisNote';

// eslint-disable-next-line react/prop-types
const MultiParagraphText = ({ value }: { value: string }) => (
  <Flex direction="column" gap="sm">
    {
      // eslint-disable-next-line react/prop-types
      value.split('\n').map(p => (
        <Text>{p}</Text>
      ))
    }
  </Flex>
);

export const BookmarkPage = () => {
  const { id } = useParams();

  const { data: itemData, isLoading: itemLoading } = useShowQuery(id || '');
  const { data: notesData, isLoading: notesLoading } = useRecentNotesQuery(
    id || '',
  );

  const isLoading = useMemo(
    () => itemLoading || notesLoading,
    [itemLoading, notesLoading],
  );
  const [uploadImage] = useUploadImageMutation();
  const [updateTitle] = useUpdateTitleMutation();
  const [updateTranscript] = useUpdateTranscriptMutation();
  const [destroyItem] = useDestroyMutation();
  const [createNote] = useCreateNoteMutation();

  const item = useMemo(() => {
    if (itemData) {
      return new Bookmark(itemData);
    }

    return null;
  }, [itemData]);

  const notes = useMemo(() => {
    if (notesData) {
      return notesData.items.map(n => new AnalysisNote(n)).reverse();
    }

    return [] as AnalysisNote[];
  }, [notesData]);

  const imageUpdated = useCallback(
    (file: File) => {
      const formData = new FormData();
      formData.append('image_file', file);
      uploadImage({ id: id || '', formData });
    },
    [id, uploadImage],
  );

  const titleUpdated = useCallback(
    (title: string) => {
      updateTitle({ id: id || '', title });
    },
    [id, updateTitle],
  );

  const transcriptUpdated = useCallback(
    (transcript: string) => {
      updateTranscript({ id: id || '', transcript });
    },
    [id, updateTranscript],
  );

  const deleteClicked = useCallback(() => {
    modals.openConfirmModal({
      title: 'Are you sure you want to delete?',
      centered: true,
      labels: { confirm: 'Yes', cancel: 'No' },
      confirmProps: { color: 'red' },
      onConfirm: () => {
        destroyItem(id || '').then(() => {
          window.location.href = '/#/';
        });
      },
    });
  }, [destroyItem, id]);

  const newNoteCreateClicked = useCallback(
    async (content: string) => {
      createNote({
        content,
        bookmarkID: id || '',
      });
    },
    [createNote, id],
  );

  return (
    <Container fluid>
      <Flex direction="row" wrap="wrap" justify="center" gap="sm">
        <Flex direction="column" gap="md" style={{ flex: 1 }}>
          <Paper p="sm" withBorder>
            <Flex wrap="wrap" justify="center">
              <Box style={{ position: 'relative' }} w={160} h={160}>
                <EditableImage
                  style={{}}
                  src={item?.image || ''}
                  w={160}
                  h={160}
                  onImageUpdated={imageUpdated}
                />
              </Box>
              <Flex direction="column" ml="sm" gap="md" style={{ flex: 1 }}>
                <Flex justify="space-between">
                  <EditableTextField
                    currentValue={item?.title || ''}
                    onValueSubmitted={titleUpdated}
                  >
                    <Title order={1}>{item?.title}</Title>
                  </EditableTextField>
                  <Menu shadow="md" width={200}>
                    <Menu.Target>
                      <ActionIcon>
                        <IconSettings size="2rem" stroke={1.5} color="gray" />
                      </ActionIcon>
                    </Menu.Target>

                    <Menu.Dropdown>
                      <Menu.Item
                        color="red"
                        leftSection={
                          <IconTrash
                            style={{ width: rem(14), height: rem(14) }}
                          />
                        }
                        onClick={deleteClicked}
                      >
                        Delete
                      </Menu.Item>
                    </Menu.Dropdown>
                  </Menu>
                </Flex>
              </Flex>
            </Flex>
          </Paper>
          <Paper p="sm" withBorder>
            <Flex direction="column" gap="md">
              <Title order={4}>Transcript</Title>
              <EditableTextField
                currentValue={item?.transcript || ''}
                multiline
                onValueSubmitted={transcriptUpdated}
              >
                <MultiParagraphText value={item?.transcript || ''} />
              </EditableTextField>
            </Flex>
          </Paper>
        </Flex>
        <Paper
          withBorder
          p="sm"
          style={{ display: 'flex', flexDirection: 'column' }}
        >
          <Flex direction="column" gap="sm" style={{ flex: 1 }}>
            <Title order={4} w={240}>
              Analysis
            </Title>
            <AnalysisPanel notes={notes} createNote={newNoteCreateClicked} />
          </Flex>
        </Paper>
      </Flex>
      <LoadingOverlay visible={isLoading} />
    </Container>
  );
};
