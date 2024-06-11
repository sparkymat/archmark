import { Card, Flex, Text } from '@mantine/core';
import React from 'react';

export interface AnalysisPanelProps {}

// eslint-disable-next-line no-empty-pattern
export const AnalysisPanel = ({}: AnalysisPanelProps) => (
  <Flex direction="column" gap="sm">
    <Card withBorder p="sm" radius={8}>
      <Text>This is a test</Text>
    </Card>
  </Flex>
);
