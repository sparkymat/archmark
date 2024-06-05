import { ActionIcon, Button, Flex, Input, Modal } from '@mantine/core';
import { IconPencil } from '@tabler/icons-react';
import React, { ChangeEvent, useCallback, useEffect, useState } from 'react';

export interface EditableFieldProps {
  currentValue: string;
  onValueSubmitted(_val: string);
}

export const EditableTextField = ({
  children,
  currentValue,
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
          <Input
            placeholder="Set new value"
            value={newValue}
            onChange={valueChanged}
          />
          <Button variant="filled" onClick={modalSubmitted}>
            Update
          </Button>
        </Flex>
      </Modal>
    </Flex>
  );
};
