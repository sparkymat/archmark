import { Button, Flex, Paper, Space, TextInput } from '@mantine/core';
import React, { ChangeEvent, useCallback } from 'react';
import { IconFilter, IconSquarePlus } from '@tabler/icons-react';

export interface FilterBarProps {
  query: string;
  onQueryChanged(_evt: ChangeEvent<HTMLInputElement>);
  filterLocation: string;
  showCreateModal();
}

export const FilterBar = ({
  query,
  onQueryChanged,
  filterLocation,
  showCreateModal,
}: FilterBarProps) => {
  const searchSubmitted = useCallback(
    (evt: React.KeyboardEvent<HTMLInputElement>) => {
      if (evt.keyCode === 13) {
        window.location.href = filterLocation;
      }
    },
    [filterLocation],
  );

  return (
    <Paper>
      <Flex p="sm">
        <TextInput
          placeholder="Search by name"
          value={query}
          onChange={onQueryChanged}
          onKeyDown={searchSubmitted}
          style={{ flex: 1 }}
        />
        <Space w="sm" />
        <Button
          variant="outline"
          leftSection={<IconFilter size={14} />}
          component="a"
          href={filterLocation}
        >
          Filter
        </Button>
        <Space w="sm" />
        <Button
          variant="filled"
          leftSection={<IconSquarePlus size={14} />}
          onClick={showCreateModal}
        >
          New
        </Button>
      </Flex>
    </Paper>
  );
};
