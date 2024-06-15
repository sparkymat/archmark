import { Button, Card, Flex, Input, Space, Text } from '@mantine/core';
import React, { ChangeEvent, useCallback, useState } from 'react';
import { AnalysisNote } from '../../models/AnalysisNote';

export interface AnalysisPanelProps {
  notes: AnalysisNote[];
  createNote(_content: string): Promise<void>;
}

// eslint-disable-next-line no-empty-pattern
export const AnalysisPanel = ({ notes, createNote }: AnalysisPanelProps) => {
  const [newNote, setNewNote] = useState<string>('');

  const newNoteUpdated = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setNewNote(evt.target.value);
  }, []);

  const newNoteCreateClicked = useCallback(() => {
    createNote(newNote).then(() => {
      setNewNote('');
    });
  }, [createNote, newNote]);

  return (
    <Flex direction="column" gap="sm" style={{ flex: 1 }}>
      <Space style={{ flex: 1 }} h="sm" />

      {notes.map(n => (
        <Card withBorder p="sm" radius={8}>
          <Text>{n.content}</Text>
        </Card>
      ))}

      <Flex gap="sm">
        <Input value={newNote} onChange={newNoteUpdated} />
        <Button onClick={newNoteCreateClicked}>Ask</Button>
      </Flex>
    </Flex>
  );
};
