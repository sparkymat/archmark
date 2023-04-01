import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { HashRouter } from 'react-router-dom';
import { ThemeProvider } from '@mui/material';

import { store } from './store';
import App from './components/App';
import theme from './theme';

const element = document.getElementById('archmark-app');

if (element) {
  ReactDOM.render(
    <Provider store={store}>
      <HashRouter>
        <ThemeProvider theme={theme}>
          <App />
        </ThemeProvider>
      </HashRouter>
    </Provider>,
    element,
  );
}
