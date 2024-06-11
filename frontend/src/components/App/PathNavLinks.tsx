import React from 'react';
import { Menu, UnstyledButton, useMantineColorScheme } from '@mantine/core';
import { IconBrightness } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';

interface Path {
  label: string;
  icon?: React.JSX.Element;
  onClick();
}

interface PathGroup {
  label: string;
  paths: Path[];
}

export const PathNavLinks = () => {
  const navigate = useNavigate();
  const { toggleColorScheme } = useMantineColorScheme();

  const pathEntries: (Path | PathGroup)[] = [
    {
      onClick: () => navigate('/'),
      label: 'home',
    },
    {
      label: 'settings',
      paths: [
        {
          label: 'toggle color',
          onClick: toggleColorScheme,
          icon: <IconBrightness strokeWidth={1} size={14} />,
        },
      ],
    },
  ];

  return (
    <>
      {pathEntries.map(pe =>
        (pe as PathGroup).paths !== undefined ? (
          <Menu shadow="md" width={200} key={pe.label}>
            <Menu.Target>
              <UnstyledButton component="a">{pe.label}</UnstyledButton>
            </Menu.Target>
            <Menu.Dropdown>
              {(pe as PathGroup).paths.map(c => (
                <Menu.Item
                  key={c.label}
                  variant="subtle"
                  tt="lowercase"
                  onClick={c.onClick}
                  leftSection={c.icon}
                >
                  {c.label}
                </Menu.Item>
              ))}
            </Menu.Dropdown>
          </Menu>
        ) : (
          <UnstyledButton
            component="a"
            key={pe.label}
            variant="subtle"
            tt="lowercase"
            onClick={(pe as Path).onClick}
          >
            {pe.label}
          </UnstyledButton>
        ),
      )}
    </>
  );
};
