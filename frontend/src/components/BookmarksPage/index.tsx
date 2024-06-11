import {
  Button,
  Card,
  Container,
  Flex,
  Group,
  Image,
  LoadingOverlay,
  Modal,
  SimpleGrid,
  Text,
  TextInput,
  Textarea,
  Title,
} from '@mantine/core';
import React, {
  ChangeEvent,
  useCallback,
  useEffect,
  useMemo,
  useState,
} from 'react';
import { useParams } from 'react-router-dom';

import { useCreateMutation, useSearchQuery } from '../../slices/Bookmark';
import { FilterBar } from '../FilterBar';
import { Pagination } from '../Pagination';
import { Bookmark } from '../../models/Bookmark';

export const BookmarksPage = () => {
  const { page: pageString, query } = useParams();

  const [newQuery, setNewQuery] = useState<string>(query || '');
  const [createTitle, setCreateTitle] = useState<string>('');
  const [createURL, setCreateURL] = useState<string>('');
  const [createDescription, setCreateDescription] = useState<string>('');
  const [createShown, setCreateShown] = useState<boolean>(false);

  const pageNumber = useMemo((): number => {
    let page = parseInt(pageString || '', 10);

    if (Number.isNaN(page)) {
      page = 1;
    }

    return page;
  }, [pageString]);

  const pageSize = 20;

  const currentFilterURL = useMemo(
    () => (query ? `/#/b/search/${query}` : '/#/b'),
    [query],
  );

  const newFilterURL = useMemo(
    () => (newQuery ? `/#/b/search/${newQuery}` : '/#/b'),
    [newQuery],
  );

  const { data: itemsData, isLoading: itemsLoading } = useSearchQuery({
    query: query || '',
    pageNumber,
    pageSize,
  });

  const [createItem, { isLoading: isCreating }] = useCreateMutation();

  const createClicked = useCallback(() => {
    createItem({
      title: createTitle,
      description: createDescription,
      url: createURL,
    }).then(res => {
      window.location.href = `/#/a/${(res as any).data.id}`;
    });
    setCreateShown(false);
  }, [createDescription, createItem, createTitle, createURL]);

  const items = useMemo(
    () =>
      itemsData?.items
        ? itemsData?.items.map(i => new Bookmark(i))
        : ([] as Bookmark[]),
    [itemsData?.items],
  );

  useEffect(() => {
    document.title = 'Bookmarks';

    if (query) {
      document.title = `Bookmarks | ${query}`;
    }
  }, [query]);

  const totalCount = useMemo(
    () => (itemsData ? itemsData.totalCount : 0),
    [itemsData],
  );

  const pageCount = useMemo(
    () => Math.ceil(totalCount / pageSize),
    [totalCount],
  );

  const newQueryChanged = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setNewQuery(evt.target.value);
  }, []);

  const createOpened = useCallback(() => {
    setCreateShown(true);
  }, []);

  const createClosed = useCallback(() => {
    setCreateShown(false);
  }, []);

  const createTitleChanged = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      setCreateTitle(evt.target.value);
    },
    [],
  );

  const createURLChanged = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setCreateURL(evt.target.value);
  }, []);

  const createDescriptionChanged = useCallback(
    (evt: ChangeEvent<HTMLTextAreaElement>) => {
      setCreateDescription(evt.target.value);
    },
    [],
  );

  return (
    <Container fluid style={{ width: '100%' }}>
      <Flex direction="column" gap="md">
        <FilterBar
          query={newQuery}
          onQueryChanged={newQueryChanged}
          filterLocation={newFilterURL}
          showCreateModal={createOpened}
        />
        {query && <Text fs="italic">{`Filtering by: ${query}`}</Text>}
        <SimpleGrid cols={{ base: 1, sm: 2, lg: 5 }}>
          {items.map(e => (
            <Card
              shadow="sm"
              padding="sm"
              withBorder
              component="a"
              href={`/#/a/${e.id}`}
            >
              <Card.Section>
                <Image src={e.image} height={160} alt={e.title} />
              </Card.Section>

              <Group justify="space-between" mt="md" mb="xs">
                <Title order={4}>{e.title}</Title>
              </Group>
            </Card>
          ))}
        </SimpleGrid>
        <Flex justify="center" mb="md">
          <Pagination
            pageNumber={pageNumber}
            pageCount={pageCount}
            filterURL={currentFilterURL}
          />
        </Flex>
      </Flex>
      <Modal title="New Bookmarks" opened={createShown} onClose={createClosed}>
        <Flex direction="column" gap="md">
          <TextInput
            placeholder="URL"
            value={createURL}
            onChange={createURLChanged}
          />
          <TextInput
            placeholder="Title"
            value={createTitle}
            onChange={createTitleChanged}
          />
          <Textarea
            placeholder="Description"
            value={createDescription}
            onChange={createDescriptionChanged}
          />
          <Button variant="filled" onClick={createClicked}>
            Add
          </Button>
        </Flex>
        <LoadingOverlay visible={isCreating} />
      </Modal>
      <LoadingOverlay visible={itemsLoading} />
    </Container>
  );
};
