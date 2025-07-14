// tailwind.config.ts
import type { Config } from 'tailwindcss'
import typography from '@tailwindcss/typography'

const config: Config = {
    darkMode: 'class',
    content: [],
    theme: {
        extend: {
            typography: () => ({
                DEFAULT: {
                    css: {
                        color: 'var(--color-foreground)',
                        backgroundColor: 'var(--color-background)',
                        '--tw-prose-body': 'var(--color-foreground)',
                        '--tw-prose-headings': 'var(--color-foreground)',
                        '--tw-prose-links': 'var(--color-primary)',
                        '--tw-prose-bold': 'var(--color-foreground)',
                        '--tw-prose-counters': 'var(--color-muted-foreground)',
                        '--tw-prose-bullets': 'var(--color-muted-foreground)',
                        '--tw-prose-hr': 'var(--color-border)',
                        '--tw-prose-quotes': 'var(--color-muted-foreground)',
                        '--tw-prose-quote-borders': 'var(--color-border)',
                        '--tw-prose-code': 'var(--color-accent-foreground)',
                        '--tw-prose-pre-code': 'var(--color-accent-foreground)',
                        '--tw-prose-pre-bg': 'var(--color-card)',
                        '--tw-prose-th-borders': 'var(--color-border)',
                        '--tw-prose-td-borders': 'var(--color-border)',
                        '--tw-prose-invert-body': 'var(--color-foreground)',
                        '--tw-prose-invert-headings': 'var(--color-foreground)',
                        '--tw-prose-invert-links': 'var(--color-primary)',
                        '--tw-prose-invert-bold': 'var(--color-foreground)',
                        '--tw-prose-invert-counters': 'var(--color-muted-foreground)',
                        '--tw-prose-invert-bullets': 'var(--color-muted-foreground)',
                        '--tw-prose-invert-hr': 'var(--color-border)',
                        '--tw-prose-invert-quotes': 'var(--color-muted-foreground)',
                        '--tw-prose-invert-quote-borders': 'var(--color-border)',
                        '--tw-prose-invert-code': 'var(--color-accent-foreground)',
                        '--tw-prose-invert-pre-code': 'var(--color-accent-foreground)',
                        '--tw-prose-invert-pre-bg': 'var(--color-card)',
                        '--tw-prose-invert-th-borders': 'var(--color-border)',
                        '--tw-prose-invert-td-borders': 'var(--color-border)',
                        maxWidth: 'unset',
                        fontFamily: [
                            '-apple-system',
                            'BlinkMacSystemFont',
                            '"Segoe UI"',
                            '"Noto Sans"',
                            'Helvetica',
                            'Arial',
                            'sans-serif',
                            '"Apple Color Emoji"',
                            '"Segoe UI Emoji"'
                        ].join(', '),
                        h1: {
                            fontSize: '2em',
                            fontWeight: '600',
                            borderBottom: '1px solid var(--color-border)',
                            paddingBottom: '0.3em',
                            margin: '0.67em 0'
                        },
                        h2: {
                            fontSize: '1.5em',
                            fontWeight: '600',
                            borderBottom: '1px solid var(--color-border)',
                            paddingBottom: '0.3em'
                        },
                        code: {
                            backgroundColor: 'var(--color-muted)',
                            borderRadius: '6px',
                            padding: '0.2em 0.4em',
                            fontSize: '85%',
                            fontWeight: 'normal'
                        },
                        pre: {
                            backgroundColor: 'var(--color-card)',
                            borderRadius: '6px',
                            padding: '1rem',
                            overflow: 'auto',
                            code: {
                                backgroundColor: 'transparent',
                                padding: 0,
                                fontSize: '85%'
                            }
                        },
                        table: {
                            borderCollapse: 'collapse',
                            width: 'max-content',
                            th: {
                                fontWeight: '600',
                                border: '1px solid var(--color-border)',
                                padding: '6px 13px',
                                backgroundColor: 'var(--color-card)'
                            },
                            td: {
                                border: '1px solid var(--color-border)',
                                padding: '6px 13px'
                            },
                            tr: {
                                backgroundColor: 'var(--color-background)',
                                borderTop: '1px solid var(--color-border)',
                                '&:nth-child(2n)': {
                                    backgroundColor: 'var(--color-muted)'
                                }
                            }
                        },
                        blockquote: {
                            borderLeft: '0.25em solid var(--color-border)',
                            color: 'var(--color-muted-foreground)',
                            padding: '0 1em',
                            margin: 0
                        }
                    }
                },
                dark: {
                    css: {
                        color: 'var(--color-foreground)',
                        backgroundColor: 'var(--color-background)',
                        '--tw-prose-body': 'var(--color-foreground)',
                        '--tw-prose-headings': 'var(--color-foreground)',
                        '--tw-prose-links': 'var(--color-primary)',
                        '--tw-prose-bold': 'var(--color-foreground)',
                        '--tw-prose-counters': 'var(--color-muted-foreground)',
                        '--tw-prose-bullets': 'var(--color-muted-foreground)',
                        '--tw-prose-hr': 'var(--color-border)',
                        '--tw-prose-quotes': 'var(--color-muted-foreground)',
                        '--tw-prose-quote-borders': 'var(--color-border)',
                        '--tw-prose-code': 'var(--color-accent-foreground)',
                        '--tw-prose-pre-code': 'var(--color-accent-foreground)',
                        '--tw-prose-pre-bg': 'var(--color-card)',
                        '--tw-prose-th-borders': 'var(--color-border)',
                        '--tw-prose-td-borders': 'var(--color-border)',
                        maxWidth: 'unset',
                        fontFamily: [
                            '-apple-system',
                            'BlinkMacSystemFont',
                            '"Segoe UI"',
                            '"Noto Sans"',
                            'Helvetica',
                            'Arial',
                            'sans-serif',
                            '"Apple Color Emoji"',
                            '"Segoe UI Emoji"'
                        ].join(', '),
                        h1: {
                            borderBottom: '1px solid var(--color-border)'
                        },
                        h2: {
                            borderBottom: '1px solid var(--color-border)'
                        },
                        code: {
                            backgroundColor: 'var(--color-muted)'
                        },
                        pre: {
                            backgroundColor: 'var(--color-card)',
                            code: {
                                color: 'var(--color-accent-foreground)'
                            }
                        },
                        table: {
                            th: {
                                border: '1px solid var(--color-border)',
                                backgroundColor: 'var(--color-card)'
                            },
                            td: {
                                border: '1px solid var(--color-border)'
                            },
                            tr: {
                                backgroundColor: 'var(--color-background)',
                                '&:nth-child(2n)': {
                                    backgroundColor: 'var(--color-muted)'
                                }
                            }
                        },
                        blockquote: {
                            borderLeftColor: 'var(--color-border)',
                            color: 'var(--color-muted-foreground)'
                        }
                    }
                }
            })
        },
    },
    plugins: [typography],
}

export default config