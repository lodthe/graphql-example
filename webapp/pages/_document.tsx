import * as React from 'react';
import Document, { Html, Head, Main, NextScript } from 'next/document';
import { useQuery } from '@apollo/client';
import QUERY_MATCHES from './graphql/queryMatches.graphql';
import ADD_COMMENT from './graphql/createComment.graphql';

export default class MyDocument extends Document {
  render() {
    const { data, loading, error } = this.props.client.query(QUERY_MATCHES);

    // check for errors
    if (error) {
      console.log('error', error);

      return <p>:( an error happened</p>;
    }

    // if all good return data
    return (
      <div className="container">
        <Head>
          <title>Mafia</title>
          <link rel='icon' href='/favicon.ico' />
        </Head>

        <h1>Matches</h1>

        {loading && <p>Loading...</p>}

        <div>
          {data.matches.map((match) => (
            <div key={match.id}>{match.createdAt}</div>
          ))}
        </div>
      </div>
    );
  }
}


MyDocument.getInitialProps = async (ctx) => {
  const originalRenderPage = ctx.renderPage;

  ctx.renderPage = () =>
      originalRenderPage({
      enhanceApp: (App) =>
          function EnhanceApp(props) {
          return <App {...props} />;
          },
      });

  const initialProps = await Document.getInitialProps(ctx);

  return {
      ...initialProps,
      styles: [
      ...React.Children.toArray(initialProps.styles),
      ],
  };
  };