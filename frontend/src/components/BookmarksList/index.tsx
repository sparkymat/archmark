import React, { ChangeEvent, useCallback } from 'react';
import {
  Anchor,
  Button,
  Container,
  Flex,
  Modal,
  Space,
  Text,
  Title,
  useMantineTheme,
} from '@mantine/core';
import moment from 'moment';
import Bookmark from '../../models/Bookmark';
import URLDisplay from '../URLDisplay';
import Paginator from '../Paginatior';

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
          <Paginator
            basePath={basePath}
            pageNumber={pageNumber}
            pageSize={pageSize}
            totalCount={totalCount}
          />
        </Flex>
      )}
      <Modal
        opened={categoryModalOpen}
        onClose={hideCategoryModal}
        title="Category modal"
      >
        <div className="uk-container-small uk-flex uk-flex-row uk-flex-between uk-modal-title">
          <h2 className="uk-modal-title">Choose category</h2>
          {/* eslint-disable-next-line max-len */}
          {/* eslint-disable-next-line jsx-a11y/anchor-has-content, jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
        </div>
        <div className="uk-flex uk-flex-row">
          <input
            type="text"
            className="uk-input"
            // eslint-disable-next-line jsx-a11y/no-autofocus
            autoFocus={categoryModalOpen}
            value={categoryModalName}
            onChange={nameChange}
            placeholder="Type new category or choose below"
          />
          <button
            type="button"
            className="uk-button uk-margin-small-left"
            onClick={categoryModalSubmitted}
          >
            Update
          </button>
        </div>
        <table className="uk-table">
          <tbody>
            {categories &&
              categories.slice(0, 5).map(c => (
                <tr className="uk-padding-bottom">
                  <td className="uk-text-large uk-padding-remove">
                    {/* eslint-disable-next-line jsx-a11y/anchor-is-valid, jsx-a11y/click-events-have-key-events, jsx-a11y/no-static-element-interactions */}
                    <a
                      className="uk-link-muted"
                      onClick={() => categoryModalNameChanged(c)}
                    >
                      {c}
                    </a>
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      </Modal>
      <Modal
        opened={deleteModalOpen}
        onClose={hideDeleteModal}
        withCloseButton={false}
        centered
      >
        <Flex direction="column" align="center" p="xl">
          <Title order={3} weight={300} my="lg">
            Are you sure?
          </Title>
          <Flex justify="space-between" my="md">
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
