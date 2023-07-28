import React, { ChangeEvent, useCallback, useMemo } from 'react';
import {
  ActionIcon,
  Anchor,
  Button,
  Container,
  Flex,
  Modal,
  Pagination,
  Space,
  Text,
  TextInput,
  Title,
  useMantineTheme,
} from '@mantine/core';
import moment from 'moment';
import { IconX } from '@tabler/icons-react';
import Bookmark from '../../models/Bookmark';
import URLDisplay from '../URLDisplay';

interface BookmarksListProps {
  basePath: string;
  bookmarks: Bookmark[];
  categories: string[];
  pageNumber: number;
  pageSize: number;
  // eslint-disable-next-line react/require-default-props
  totalCount: number;
  allowCategoryChange: boolean;
  categoryModalOpen: boolean;
  categoryModalName: string;
  showArchiveButton: boolean;
  showUnarchiveButton: boolean;
  deleteModalOpen: boolean;
  // eslint-disable-next-line react/no-unused-prop-types, react/require-default-props
  categoryModalBookmarkID?: string;
  hideCategoryModal(): void;
  showCategoryModal(_bookmarkID: string): void;
  categoryModalNameChanged(_val: string): void;
  categoryModalSubmitted(): void;
  hideDeleteModal(): void;
  showDeleteModal(_bookmarkID: string): void;
  deleteModalSubmitted(): void;
  unarchiveClicked(_bookmarkID: string): void;
}

const BookmarksList = ({
  basePath,
  bookmarks,
  categories,
  pageNumber,
  pageSize,
  totalCount,
  allowCategoryChange,
  categoryModalOpen,
  deleteModalOpen,
  hideCategoryModal,
  showCategoryModal,
  hideDeleteModal,
  showDeleteModal,
  categoryModalName,
  categoryModalNameChanged,
  categoryModalSubmitted,
  showArchiveButton,
  showUnarchiveButton,
  deleteModalSubmitted,
  unarchiveClicked,
}: BookmarksListProps) => {
  const nameChange = useCallback(
    (evt: ChangeEvent<HTMLInputElement>) => {
      categoryModalNameChanged(evt.target.value);
    },
    [categoryModalNameChanged],
  );
  const theme = useMantineTheme();

  const pageCount = useMemo(
    () => Math.ceil(totalCount / pageSize),
    [pageSize, totalCount],
  );

  return (
    <Container>
      {(!bookmarks || bookmarks.length === 0) && (
        <Flex p="xl" justify="center" sx={{ border: '1px dashed #999999' }}>
          <Text>No bookmarks found</Text>
        </Flex>
      )}
      {bookmarks &&
        bookmarks.map((b, i) => (
          <Flex direction="column" mb="lg">
            <Anchor
              underline={false}
              color={theme.colorScheme === 'dark' ? '#dddddd' : '#444444'}
              href={b.url}
              target="_blank"
              rel="noreferrer"
            >
              <Flex align="center">
                <Text size="sm">{(pageNumber - 1) * pageSize + i + 1}.</Text>
                <Space w="sm" />
                {b.title ? (
                  <Title order={3} size={24} weight={300}>
                    {b.title}
                  </Title>
                ) : (
                  <URLDisplay colorScheme={theme.colorScheme} value={b.url} />
                )}
              </Flex>
            </Anchor>
            <Flex>
              <Text>added {moment(b.created_at).toNow(true)} ago</Text>
              <Text mx="xs">⚬</Text>
              {allowCategoryChange && (
                // eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions
                <Anchor
                  underline={false}
                  onClick={() => showCategoryModal(b.id)}
                >
                  <Text>{b.category ? b.category : 'Uncategorized'}</Text>
                </Anchor>
              )}
              {!allowCategoryChange && (
                <Text>{b.category ? b.category : 'Uncategorized'}</Text>
              )}
              {b.file_path && (
                <>
                  <Text mx="xs">⚬</Text>
                  <Anchor
                    underline={false}
                    href={`/uploads/${b.file_path}`}
                    target="_blank"
                    rel="noreferrer"
                  >
                    Cached
                  </Anchor>
                </>
              )}
              {showArchiveButton && (
                <>
                  <Text mx="xs">⚬</Text>
                  {/* eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
                  <Anchor
                    underline={false}
                    onClick={() => showDeleteModal(b.id)}
                  >
                    Delete
                  </Anchor>
                </>
              )}
              {showUnarchiveButton && (
                <>
                  <Text mx="xs">⚬</Text>
                  {/* eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
                  <Anchor
                    underline={false}
                    onClick={() => unarchiveClicked(b.id)}
                  >
                    Restore
                  </Anchor>
                </>
              )}
            </Flex>
          </Flex>
        ))}
      {bookmarks && bookmarks.length > 0 && (
        <Flex justify="center" mb="md">
          <Pagination
            value={pageNumber}
            boundaries={1}
            total={pageCount}
            position="center"
            withEdges
            getItemProps={p => ({
              component: 'a',
              href: `${basePath}page/${p}`,
            })}
            getControlProps={control => {
              if (control === 'first') {
                return { component: 'a', href: `${basePath}page/1` };
              }

              if (control === 'last') {
                return {
                  component: 'a',
                  href: `${basePath}page/${pageCount}`,
                };
              }

              if (control === 'next') {
                return {
                  component: 'a',
                  href:
                    pageNumber + 1 > pageCount
                      ? `${basePath}page/${pageNumber}`
                      : `${basePath}page/${pageNumber + 1}`,
                };
              }

              if (control === 'previous') {
                return {
                  component: 'a',
                  href:
                    pageNumber - 1 < 1
                      ? `${basePath}page/${pageNumber}`
                      : `${basePath}page/${pageNumber - 1}`,
                };
              }

              return {};
            }}
          />
        </Flex>
      )}
      <Modal
        opened={categoryModalOpen}
        onClose={hideCategoryModal}
        withCloseButton={false}
      >
        <Flex direction="column" align="stretch" p="lg">
          <Title align="center" order={3} weight={300} mb="xl">
            Change category
          </Title>
          <Flex>
            <TextInput
              sx={{ flex: 1 }}
              autoFocus={categoryModalOpen}
              value={categoryModalName}
              onChange={nameChange}
              placeholder="Type new category or choose below"
              rightSection={
                <ActionIcon onClick={() => categoryModalNameChanged('')}>
                  <IconX size="1.125rem" />
                </ActionIcon>
              }
            />
            <Button variant="outline" ml="sm" onClick={categoryModalSubmitted}>
              Update
            </Button>
          </Flex>
          <Flex direction="column" justify="stretch">
            {categories &&
              categories.slice(0, 5).map(c => (
                <Button
                  variant="subtle"
                  onClick={() => categoryModalNameChanged(c)}
                  mt="xs"
                >
                  {c}
                </Button>
              ))}
          </Flex>
        </Flex>
      </Modal>
      <Modal
        opened={deleteModalOpen}
        onClose={hideDeleteModal}
        withCloseButton={false}
      >
        <Flex direction="column" align="center" p="md">
          <Title order={3} weight={300} mb="xl">
            Are you sure?
          </Title>
          <Flex justify="space-between">
            <Button
              color="red"
              variant="outline"
              onClick={() => hideDeleteModal()}
            >
              No
            </Button>

            <Space w="xl" />

            <Button
              color="green"
              variant="outline"
              onClick={() => deleteModalSubmitted()}
            >
              Yes
            </Button>
          </Flex>
        </Flex>
      </Modal>
    </Container>
  );
};

export default BookmarksList;
