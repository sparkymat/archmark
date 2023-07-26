import { Flex, Title } from '@mantine/core';
import React from 'react';

interface URLDisplayProps {
  value: string;
  colorScheme: 'light' | 'dark';
}

const URLDisplay = ({ value, colorScheme }: URLDisplayProps) => {
  const regex = /(http|https):\/\/([^/]*)\/(.*)/;
  const matches = value.match(regex);

  if (!matches) {
    return <span>{value}</span>;
  }

  const protocol = matches[1];
  const host = matches[2];
  const path = matches[3];

  return (
    <Flex>
      <Title
        order={6}
        weight={300}
        size="lg"
        color={colorScheme === 'dark' ? '#666666' : '#999999'}
      >
        {protocol}://
      </Title>
      <Title
        order={6}
        size="lg"
        weight={400}
        color={colorScheme === 'dark' ? '#cccccc' : '#222222'}
      >
        {host}
      </Title>
      <Title
        order={6}
        weight={300}
        size="lg"
        color={colorScheme === 'dark' ? '#666666' : '#999999'}
      >
        /{path}
      </Title>
    </Flex>
  );
};

export default URLDisplay;
