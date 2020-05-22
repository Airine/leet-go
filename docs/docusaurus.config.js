module.exports = {
  title: 'Leet Go',
  tagline: '用Golang刷LeetCode!', // one sentence description
  url: 'https://blog.aaron-xin.tech/',
  baseUrl: '/leet-go/',
  favicon: 'img/favicon.ico',
  organizationName: 'Airine', // Usually your GitHub org/user name.
  projectName: 'leet-go', // Usually your repo name.
  themeConfig: {
    navbar: {
      title: 'Leet Go',
      logo: {
        alt: 'SUSTech CANStudio Logo',
        src: 'img/logo.svg',
      },
      links: [
        {
          to: 'docs/doc1',
          activeBasePath: 'docs',
          label: 'Solutions',
          position: 'left',
        },
        {href: 'https://airine.github.io/blog', label: 'Blog', position: 'left'},
        {href: 'https://airine.github.io/', label: 'Home', position: 'right'},
        {
          href: 'https://github.com/airine/leet-go',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Solutions',
          items: [
            {
              label: 'TODO List',
              to: 'docs/doc1',
            }
          ],
        },
        {
          title: 'Social',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/airine',
            }
          ],
        },
        {
          title: 'Links',
          items: [
            {
              label: 'Samuel',
              href: 'https://hustergs.github.io'
            }
          ]
        }
      ],
      copyright: `Copyright © ${new Date().getFullYear()} LeetGo @ 一口闰心`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl:
            'https://github.com/airine/leet-go/edit/master/docs/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
