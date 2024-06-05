import React, { useCallback } from 'react';
import { NavLink } from '@mantine/core';
import { IconBooks, IconDashboard, IconSettings } from '@tabler/icons-react';
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';

import { selectApp } from '../../store/selectors';
import { AppDispatch } from '../../store';
import { updatePath } from '../../slices/App';

interface Path {
  href: string;
  label: string;
  icon?: React.JSX.Element;
}

interface PathGroup {
  label: string;
  icon: React.JSX.Element;
  paths: Path[];
}

export interface PathNavLinksProps {
  toggleNav(): void;
}

export const PathNavLinks = ({ toggleNav }: PathNavLinksProps) => {
  const dispatch = useDispatch<AppDispatch>();
  const { path } = useSelector(selectApp);
  const navigate = useNavigate();

  const onNavClick = useCallback(
    (p: string) => {
      dispatch(updatePath(p));
      navigate(p);
      toggleNav();
    },
    [dispatch, navigate, toggleNav],
  );

  const isActive = useCallback(
    (newPath: string): boolean => newPath === path,
    [path],
  );

  const pathEntries: (Path | PathGroup)[] = [
    {
      href: '/',
      label: 'Dashboard',
      icon: <IconDashboard size="1rem" stroke={1.5} />,
    },
  ];

  return (
    <>
      {pathEntries.map(pe =>
        (pe as PathGroup).paths !== undefined ? (
          <NavLink
            key={pe.label}
            label={pe.label}
            variant="subtle"
            tt="uppercase"
            leftSection={(pe as PathGroup).icon}
            childrenOffset={20}
          >
            {(pe as PathGroup).paths.map(c => (
              <NavLink
                key={c.href}
                label={c.label}
                active={isActive(c.href)}
                variant="subtle"
                tt="uppercase"
                onClick={() => onNavClick(c.href)}
              />
            ))}
          </NavLink>
        ) : (
          <NavLink
            key={pe.label}
            label={pe.label}
            active={isActive((pe as Path).href)}
            variant="subtle"
            tt="uppercase"
            leftSection={(pe as Path).icon}
            onClick={() => onNavClick((pe as Path).href)}
          />
        ),
      )}
    </>
  );
};
