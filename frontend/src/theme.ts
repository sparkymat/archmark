import { createTheme } from '@mui/material';

declare module '@mui/material' {
  // eslint-disable-next-line no-unused-vars
  interface ThemeOptions {}
}

const theme = createTheme({
  palette: {
    primary: {
      main: '#0050c8',
    },
  },
  components: {
    MuiAppBar: {
      styleOverrides: {
        root: {
          boxShadow: 'none',
        },
      },
    },
    MuiOutlinedInput: {
      styleOverrides: {
        root: {
          borderRadius: 0,
          backgroundColor: '#f0f0f0',
        },
      },
    },
    MuiDrawer: {
      styleOverrides: {
        paper: {
          backgroundColor: '#003585',
        },
      },
    },
  },
});

export default theme;
