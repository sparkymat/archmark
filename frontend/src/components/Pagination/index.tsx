import React from 'react';
import { Pagination as PaginationControl } from '@mantine/core';

export interface PaginationParams {
  pageNumber: number;
  pageCount: number;
  filterURL: string;
}

export const Pagination = ({
  pageCount,
  pageNumber,
  filterURL,
}: PaginationParams) => (
  <PaginationControl
    value={pageNumber}
    boundaries={2}
    total={pageCount}
    withEdges
    getItemProps={p => ({
      component: 'a',
      href: `${filterURL}/p/${p}`,
    })}
    getControlProps={control => {
      if (control === 'first') {
        return { component: 'a', href: filterURL };
      }

      if (control === 'last') {
        return {
          component: 'a',
          href: `${filterURL}/p/${pageCount}`,
        };
      }

      if (control === 'next') {
        return {
          component: 'a',
          href:
            pageNumber + 1 > pageCount
              ? `${filterURL}/p/${pageNumber}`
              : `${filterURL}/p/${pageNumber + 1}`,
        };
      }

      if (control === 'previous') {
        return {
          component: 'a',
          href:
            pageNumber - 1 < 1
              ? `${filterURL}/p/${pageNumber}`
              : `${filterURL}/p/${pageNumber - 1}`,
        };
      }

      return {};
    }}
  />
);
