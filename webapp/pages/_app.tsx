import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client';
import { ApolloProvider } from '@apollo/client';
import React, { Component } from "react";
import '../styles/globals.css';
import styles from '../styles/Mafia.module.css';
import { ThemeProvider } from '@mui/material/styles';
import theme from '../styles/theme/lightThemeOptions';

const client = new ApolloClient({
  uri: 'http://localhost:8080/query',
  cache: new InMemoryCache(),
});

const MyApp = (props) => {
  const { Component, pageProps } = props;

  return (
    <React.Fragment>
      <ApolloProvider client={client}>
        <ThemeProvider theme={theme}>
          <Component {...pageProps} client={client} />
        </ThemeProvider>
      </ApolloProvider>
    </React.Fragment>
  );
}

export default MyApp;
