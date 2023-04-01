/* eslint-disable react/no-unknown-property */
import {
  Box,
  CssBaseline,
  Drawer,
  IconButton,
  List,
  ListItem,
  ListItemButton,
  ListItemIcon,
  ListItemIconProps,
  ListItemProps,
  ListItemText,
  styled,
} from '@mui/material';
import HomeIcon from '@mui/icons-material/Home';
import MenuIcon from '@mui/icons-material/Menu';
import React, { useCallback, useState } from 'react';
import { Routes, Route } from 'react-router-dom';

import Home from '../Home';

const drawerWidth = 240;

const App = () => {
  const [mobileOpen, setMobileOpen] = useState<boolean>(false);

  const handleDrawerToggle = useCallback(() => {
    setMobileOpen(!mobileOpen);
  }, [mobileOpen]);

  const navigateTo = useCallback((path: string) => {
    window.location.href = `/#${path}`;
  }, []);

  const DarkListItemIcon = styled(ListItemIcon)<ListItemIconProps>(() => ({
    color: 'white',
  }));

  const DarkListItem = styled(ListItem)<ListItemProps>(() => ({
    color: 'white',
  }));

  const drawer = (
    <div>
      <Box paddingY={4} sx={{ display: 'flex', justifyContent: 'center' }}>
        <a href="/#/">archmark</a>
      </Box>
      <List>
        <DarkListItem disablePadding onClick={() => navigateTo('/')}>
          <ListItemButton>
            <DarkListItemIcon>
              <HomeIcon />
            </DarkListItemIcon>
            <ListItemText primary="Home" />
          </ListItemButton>
        </DarkListItem>
      </List>
    </div>
  );

  const container =
    window !== undefined ? () => window.document.body : undefined;

  return (
    <div>
      <Box sx={{ display: 'flex' }}>
        <CssBaseline />
        <Box
          component="nav"
          sx={{ width: { sm: drawerWidth }, flexShrink: { sm: 0 } }}
          aria-label="mailbox folders"
        >
          {/* The implementation can be swapped with js to avoid SEO duplication of links. */}
          <Drawer
            container={container}
            variant="temporary"
            open={mobileOpen}
            onClose={handleDrawerToggle}
            ModalProps={{
              keepMounted: true, // Better open performance on mobile.
            }}
            sx={{
              display: { xs: 'block', sm: 'none' },
              '& .MuiDrawer-paper': {
                boxSizing: 'border-box',
                width: drawerWidth,
              },
            }}
          >
            {drawer}
          </Drawer>
          <Drawer
            variant="permanent"
            sx={{
              display: { xs: 'none', sm: 'block' },
              '& .MuiDrawer-paper': {
                boxSizing: 'border-box',
                width: drawerWidth,
              },
            }}
            open
          >
            {drawer}
          </Drawer>
        </Box>
        <Box
          component="main"
          sx={{
            flexGrow: 1,
            p: 3,
            width: { sm: `calc(100% - ${drawerWidth}px)` },
          }}
        >
          <Box sx={{ ml: 1, mt: 1, display: { sm: 'none' } }}>
            <IconButton
              color="inherit"
              aria-label="open drawer"
              edge="start"
              onClick={handleDrawerToggle}
            >
              <MenuIcon />
            </IconButton>
          </Box>
          <Routes>
            <Route index element={<Home />} />
          </Routes>
        </Box>
      </Box>
    </div>
  );
};
export default App;
