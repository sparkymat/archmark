import {
  ActionIcon,
  Button,
  Flex,
  Input,
  Modal,
  Textarea,
} from '@mantine/core';
import { IconPencil } from '@tabler/icons-react';
import React, { ChangeEvent, useCallback, useEffect, useState } from 'react';

export interface EditableFieldProps {
  currentValue: string;
  // eslint-disable-next-line react/require-default-props
  multiline?: boolean;
  onValueSubmitted(_val: string);
}

export const EditableTextField = ({
  children,
  currentValue,
  multiline,
  onValueSubmitted,
}: React.PropsWithChildren<EditableFieldProps>) => {
  const [modalOpen, setModalOpen] = useState<boolean>(false);
  const [newValue, setNewValue] = useState<string>(currentValue);

  useEffect(() => {
    setNewValue(currentValue);
  }, [currentValue]);

  const valueChanged = useCallback((evt: ChangeEvent<HTMLInputElement>) => {
    setNewValue(evt.target.value);
  }, []);

  const multilineValueChanged = useCallback(
    (evt: ChangeEvent<HTMLTextAreaElement>) => {
      setNewValue(evt.target.value);
    },
    [],
  );

  const modalOpened = useCallback(() => {
    setModalOpen(true);
  }, []);

  const modalClosed = useCallback(() => {
    setModalOpen(false);
  }, []);

  const modalSubmitted = useCallback(() => {
    onValueSubmitted(newValue);
    setModalOpen(false);
  }, [newValue, onValueSubmitted]);

  return (
    <Flex align="center" gap="md">
      <>
        {children}
        <ActionIcon onClick={modalOpened} color="#666666">
          <IconPencil stroke={1} size="1.25rem" />
        </ActionIcon>
      </>
      <Modal title="Edit" opened={modalOpen} onClose={modalClosed}>
        <Flex direction="column" gap="md">
          {multiline ? (
            <Textarea
              placeholder="Set new value"
              value={newValue}
              resize="vertical"
              rows={12}
              onChange={multilineValueChanged}
            />
          ) : (
            <Input
              placeholder="Set new value"
              value={newValue}
              onChange={valueChanged}
            />
          )}

          <Button variant="filled" onClick={modalSubmitted}>
            Update
          </Button>
        </Flex>
      </Modal>
    </Flex>
  );
};
