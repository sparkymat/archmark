import React from 'react';

interface PaginatorProps {
  pageNumber: number;
  pageSize: number;
  totalCount: number;
}

const Paginator = ({ pageNumber, pageSize, totalCount }: PaginatorProps) => {
  const totalPages = Math.ceil(totalCount / pageSize);

  let middlePages = [pageNumber - 1, pageNumber, pageNumber + 1];
  middlePages = middlePages.filter(p => p >= 1 && p <= totalPages);

  const pages: (number | undefined)[] = [];
  if (middlePages[0] !== 1) {
    pages.push(1);

    if (middlePages[0] !== 2) {
      pages.push(undefined);
    }
  }

  pages.push(...middlePages);

  if (middlePages[middlePages.length - 1] !== totalPages) {
    if (middlePages[middlePages.length - 1] !== totalPages - 1) {
      pages.push(undefined);
    }
    pages.push(totalPages);
  }

  const liClass = (p: number | undefined) => {
    if (!p) {
      return 'uk-disabled';
    }

    if (p === pageNumber) {
      return 'uk-active;';
    }

    return '';
  };

  return (
    <ul className="uk-pagination">
      {pageNumber !== 1 && (
        <li>
          <a href={pageNumber === 2 ? '/' : `/#/page/${pageNumber - 1}`}>
            {/* eslint-disable-next-line react/no-unknown-property */}
            <span uk-pagination-previous="" />
          </a>
        </li>
      )}
      {pages.map(p => (
        <li className={liClass(p)}>
          {p && p !== pageNumber && (
            <a href={p === 1 ? '/' : `/#/page/${p}`}>{p}</a>
          )}
          {p && p === pageNumber && <span>{p}</span>}
          {!p && <span>...</span>}
        </li>
      ))}
      {pageNumber !== totalPages && (
        <li>
          <a href={`/#/page/${pageNumber + 1}`}>
            {/* eslint-disable-next-line react/no-unknown-property */}
            <span uk-pagination-next="" />
          </a>
        </li>
      )}
    </ul>
  );
};

export default Paginator;
