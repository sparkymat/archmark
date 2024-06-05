import { ActionIcon, Button, Flex, Modal } from '@mantine/core';
import { DatePicker } from '@mantine/dates';
import { IconPencil } from '@tabler/icons-react';
import dayjs from 'dayjs';
import React, { useCallback, useEffect, useState } from 'react';

export interface EditableFieldProps {
  currentValue: dayjs.Dayjs | undefined;
  onValueSubmitted(_val: dayjs.Dayjs | undefined);
}

export const EditableDateField = ({
  children,
  currentValue,
  onValueSubmitted,
}: React.PropsWithChildren<EditableFieldProps>) => {
  const [modalOpen, setModalOpen] = useState<boolean>(false);
  const [newValue, setNewValue] = useState<Date | null>(
    currentValue ? currentValue.toDate() : null,
  );

  useEffect(() => {
    setNewValue(currentValue ? currentValue.toDate() : null);
  }, [currentValue]);

  const valueChanged = useCallback((val: Date | null) => {
    setNewValue(val);
  }, []);

  const modalOpened = useCallback(() => {
    setModalOpen(true);
  }, []);

  const modalClosed = useCallback(() => {
    setModalOpen(false);
  }, []);

  const modalSubmitted = useCallback(() => {
    if (newValue) {
      const d = dayjs(newValue);
      onValueSubmitted(d);
    } else {
      onValueSubmitted(undefined);
    }
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
        <Flex direction="column" gap="md" align="center">
          <DatePicker value={newValue} onChange={valueChanged} />
          <Button variant="filled" onClick={modalSubmitted}>
            Update
          </Button>
        </Flex>
      </Modal>
    </Flex>
  );
};
