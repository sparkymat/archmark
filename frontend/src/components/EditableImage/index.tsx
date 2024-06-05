import {
  ActionIcon,
  Box,
  FileButton,
  Flex,
  Image,
  Overlay,
  Space,
} from '@mantine/core';
import { IconCheck, IconUpload, IconX } from '@tabler/icons-react';
import React, { useCallback, useMemo, useState } from 'react';

export interface EditableImageProps {
  style: any;
  src: string;
  w: number;
  h: number;
  onImageUpdated(_f: File);
}

export const EditableImage = ({
  style,
  src,
  w,
  h,
  onImageUpdated,
}: EditableImageProps) => {
  const [file, setFile] = useState<File | null>(null);

  const fileURL = useMemo(
    (): string | null => (file ? URL.createObjectURL(file) : null),
    [file],
  );

  const avatarSubmitted = useCallback(() => {
    if (file) {
      onImageUpdated(file);
      setFile(null);
    }
  }, [file, onImageUpdated]);

  const imageCleared = useCallback(() => {
    setFile(null);
  }, []);

  return (
    <Box style={[style, { position: 'relative' }]}>
      <Overlay
        style={
          file
            ? { backgroundColor: 'rgba(255,255,255,0.3)' }
            : { backgroundColor: 'rgba(0,0,0,0)' }
        }
      >
        <Flex direction="column-reverse" align="center" h="100%" p="sm">
          <Flex justify="space-around" w="100%">
            {!file && (
              <>
                <Space w="sm" />
                <FileButton onChange={setFile} accept="image/png,image/jpeg">
                  {props => (
                    <ActionIcon
                      color="black"
                      size="lg"
                      variant="subtle"
                      // eslint-disable-next-line react/jsx-props-no-spreading
                      {...props}
                    >
                      <IconUpload size="2rem" />
                    </ActionIcon>
                  )}
                </FileButton>
                <Space w="sm" />
              </>
            )}
            {file && (
              <>
                <ActionIcon
                  color="red"
                  size="lg"
                  variant="subtle"
                  onClick={imageCleared}
                >
                  <IconX size="2rem" />
                </ActionIcon>
                <ActionIcon
                  color="green"
                  size="lg"
                  variant="subtle"
                  onClick={avatarSubmitted}
                >
                  <IconCheck size="2rem" />
                </ActionIcon>
              </>
            )}
          </Flex>
        </Flex>
      </Overlay>
      <Image w={w} h={h} fit="cover" src={fileURL || src} />
    </Box>
  );
};
