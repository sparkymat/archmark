import React, { useCallback } from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { HashRouter } from 'react-router-dom';
import {
  ColorScheme,
  ColorSchemeProvider,
  MantineProvider,
} from '@mantine/core';
import { useLocalStorage } from '@mantine/hooks';
import { store } from './store';
import App from './components/App';

const StyledApp = () => {
  const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
    key: 'mantine-color-scheme',
    defaultValue: 'light',
    getInitialValueInEffect: true,
  });

  const toggleColorScheme = useCallback(
    (value?: ColorScheme) =>
      setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark')),
    [colorScheme, setColorScheme],
  );

  return (
    <ColorSchemeProvider
      colorScheme={colorScheme}
      toggleColorScheme={toggleColorScheme}
    >
      <MantineProvider
        withGlobalStyles
        withNormalizeCSS
        theme={{
          colorScheme,
          fontFamily: 'PT Sans, sans-serif',
          headings: { fontFamily: 'Montserrat, sans-serif' },
        }}
      >
        <App />
      </MantineProvider>
    </ColorSchemeProvider>
  );
};

const element = document.getElementById('archmark-app');

if (element) {
  ReactDOM.render(
    <Provider store={store}>
      <HashRouter>
        <StyledApp />
      </HashRouter>
    </Provider>,
    element,
  );
}
