import React, { useCallback, useState } from 'react';
import { Routes, Route } from 'react-router-dom';
import {
  IconBrightness,
  IconBuildingSkyscraper,
  IconDashboard,
  IconMapPinFilled,
  IconTimelineEvent,
  IconUsers,
} from '@tabler/icons-react';
import {
  ActionIcon,
  Anchor,
  AppShell,
  Burger,
  ColorScheme,
  Container,
  Footer,
  Header,
  MediaQuery,
  NavLink,
  Navbar,
  Space,
  Title,
  useMantineTheme,
} from '@mantine/core';
import { useDispatch, useSelector } from 'react-redux';
import { useLocalStorage } from '@mantine/hooks';
import NewBookmark from '../NewBookmark';
import Search from '../Search';
import ArchivedList from '../ArchivedList';
import { updatePath } from '../../features/App/slice';
import { selectPath } from '../../features/App/selects';
import { AppDispatch } from '../../store';

interface Path {
  href: string;
  label: string;
  icon: React.JSX.Element;
}

const App = () => {
  const dispatch = useDispatch<AppDispatch>();
  const theme = useMantineTheme();
  const [opened, setOpened] = useState(false);

  const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
    key: 'mantine-color-scheme',
    defaultValue: 'light',
    getInitialValueInEffect: true,
  });

  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

  const currentPath = useSelector(selectPath);

  const paths: Path[] = [
    {
      href: '/',
      label: 'Dashboard',
      icon: <IconDashboard size="1rem" stroke={1.5} />,
    },
    {
      href: '/contacts/page/1',
      label: 'Contacts',
      icon: <IconUsers size="1rem" stroke={1.5} />,
    },
    {
      href: '/events',
      label: 'Events',
      icon: <IconTimelineEvent size="1rem" stroke={1.5} />,
    },
    {
      href: '/organisations/page/1',
      label: 'Organisations',
      icon: <IconBuildingSkyscraper size="1rem" stroke={1.5} />,
    },
    {
      href: '/locations/page/1',
      label: 'Locations',
      icon: <IconMapPinFilled size="1rem" stroke={1.5} />,
    },
  ];

  const onNavClick = useCallback(
    (p: string) => {
      dispatch(updatePath(p));
      window.location.href = `/#${p}`;
    },
    [dispatch],
  );

  const isActive = useCallback(
    (path: string): boolean => currentPath === path,
    [currentPath],
  );

  return (
    <AppShell
      styles={{
        main: {
          background:
            theme.colorScheme === 'dark'
              ? theme.colors.dark[8]
              : theme.colors.gray[0],
        },
      }}
      header={
        <Header height={{ base: 50, md: 70 }} p="md">
          <div
            style={{ display: 'flex', alignItems: 'center', height: '100%' }}
          >
            <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
              <Burger
                opened={opened}
                onClick={() => setOpened(o => !o)}
                size="sm"
                color={theme.colors.gray[6]}
                mr="xl"
              />
            </MediaQuery>

            <Anchor href="/" underline={false}>
              <Title order={3} weight={300}>
                archmark
              </Title>
            </Anchor>
            <Space w="xl" />
            <Anchor href="/#/bookmark" underline={false}>
              <Title order={6} weight={300}>
                new bookmark
              </Title>
            </Anchor>
            <Space w="xl" />
            <Anchor href="/#/deleted" underline={false}>
              <Title order={6} weight={300}>
                deleted
              </Title>
            </Anchor>
            <Space sx={{ flex: 1 }} />
            <ActionIcon onClick={() => toggleColorScheme()}>
              <IconBrightness size="1.125rem" />
            </ActionIcon>
          </div>
        </Header>
      }
    >
      <Container>
        <Routes>
          <Route index element={<Search />} />
          <Route path="/page/:page" element={<Search />} />
          <Route path="/bookmark" element={<NewBookmark />} />
          <Route path="/search" element={<Search />} />
          <Route path="/search/:query" element={<Search />} />
          <Route path="/search/:query/page/:page" element={<Search />} />
          <Route path="/deleted" element={<ArchivedList />} />
          <Route path="/deleted/page/:page" element={<ArchivedList />} />
        </Routes>
      </Container>
    </AppShell>
  );
};

export default App;
