import React, { StrictMode } from 'react';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime'; // import plugin
import utc from 'dayjs/plugin/utc'; // import plugin
import { createRoot } from 'react-dom/client';
import { HashRouter } from 'react-router-dom';
import { MantineProvider } from '@mantine/core';
import { ModalsProvider } from '@mantine/modals';
import { Notifications } from '@mantine/notifications';
import { Provider } from 'react-redux';

import { store } from './store';
import { App } from './components/App';

// eslint-disable-next-line import/extensions, import/no-unresolved
import '@mantine/core/styles.css';
// eslint-disable-next-line import/extensions, import/no-unresolved
import '@mantine/dates/styles.css';
// eslint-disable-next-line import/extensions, import/no-unresolved
import '@mantine/carousel/styles.css';

import { theme } from './theme';

dayjs.extend(relativeTime); // use plugin
dayjs.extend(utc);

const StyledApp = () => (
  <MantineProvider defaultColorScheme="auto" theme={theme}>
    <ModalsProvider>
      <Notifications position="bottom-right" />
      <App />
    </ModalsProvider>
  </MantineProvider>
);

const element = document.getElementById('archmark-app');

if (element) {
  const root = createRoot(element);
  root.render(
    <Provider store={store}>
      <HashRouter>
        <StrictMode>
          <StyledApp />
        </StrictMode>
      </HashRouter>
    </Provider>,
  );
}
