module.exports = {
    title: 'Arr.ai',
    tagline: 'The ultimate data engine',
    url: 'https://arr.ai',
    baseUrl: '/',
    onBrokenLinks: 'throw',
    favicon: 'img/favicon.svg',
    organizationName: 'arr-ai',
    projectName: 'arrai',
    themeConfig: {
        navbar: {
            title: 'Arr.ai',
            logo: {
                alt: 'Arr.ai Logo',
                src: 'img/favicon.svg',
            },
            items: [
                {
                    to: 'docs/',
                    activeBasePath: 'docs',
                    label: 'Docs',
                    position: 'left',
                },
                {
                    href: 'https://github.com/arr-ai/arrai',
                    label: 'GitHub',
                    position: 'right',
                },
            ],
        },
        googleAnalytics: {
            trackingID: "G-HKGSHFG2D0", // GA account owner: sysl.usr@gmail.com
            anonymizeIP: true,
        },
    },
    presets: [
        [
            '@docusaurus/preset-classic',
            {
                docs: {
                    sidebarPath: require.resolve('./sidebars.js'),
                    editUrl: 'https://github.com/arr-ai/arrai/edit/master/docs/',
                },
                theme: {
                    customCss: require.resolve('./src/css/custom.css'),
                },
            },
        ],
    ],
};