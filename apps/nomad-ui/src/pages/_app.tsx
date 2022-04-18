import { CacheProvider, EmotionCache } from '@emotion/react';
import CssBaseline from '@mui/material/CssBaseline';
import { ThemeProvider } from '@mui/material/styles';
import { AppProps } from 'next/app';
import Head from 'next/head';

import { createEmotionCache } from '../common/create-emotion-cache';
import { theme } from '../theme/theme';

const clientSideEmotionCache = createEmotionCache();

function CustomApp({
    Component,
    emotionCache = clientSideEmotionCache,
    pageProps,
}: AppProps & { emotionCache?: EmotionCache }) {
    return (
        <CacheProvider value={emotionCache}>
            <Head>
                <Head>
                    <meta name="viewport" content="initial-scale=1, width=device-width" />
                </Head>
            </Head>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                <Component {...pageProps} />
            </ThemeProvider>
        </CacheProvider>
    );
}

export default CustomApp;
