import React from 'react';

interface URLDisplayProps {
  value: string;
}

const URLDisplay = ({ value }: URLDisplayProps) => {
  const regex = /(http|https):\/\/([^/]*)\/(.*)/;
  const matches = value.match(regex);

  if (!matches) {
    return <span>{value}</span>;
  }

  const protocol = matches[1];
  const host = matches[2];
  const path = matches[3];

  return (
    <>
      <span className="uk-text-muted uk-text-default">{protocol}://</span>
      <span style={{ fontWeight: 400 }} className="uk-text-default">
        {host}
      </span>
      <span className="uk-text-muted uk-text-default">/{path}</span>
    </>
  );
};

export default URLDisplay;
