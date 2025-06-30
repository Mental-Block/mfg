import React from 'react';
import { useNavigate, Link } from 'react-router';
import { Button, Dropdown, MenuItemProps, MenuProps, theme as antdTheme } from 'antd';
import { MoonOutlined, SunOutlined, UserOutlined, LogoutOutlined, AppstoreOutlined } from '@ant-design/icons';

import { useWindowResize } from 'src/hooks/useWindowResize';
import { useThemeStore } from 'src/store/useThemeStore';
import { useLogoutMutation } from 'src/features/authentication/hooks/useLogoutMutation';
import { useUserStore } from 'src/store/useUserStore';
import { genericServerErrorHandler } from 'src/utils/error';
import { ItemType, MenuItemType } from 'antd/es/menu/interface';

type MenuItem = Required<MenuProps>['items'][number];

export const useHeaderMenu = (): MenuItem[] | undefined => {
  const { token } = antdTheme.useToken();
  const navigate = useNavigate();
  const theme = useThemeStore();
  const mutationLogout = useLogoutMutation();
  const [isPassed] = useWindowResize('md', false);

  const logout = useUserStore((state) => state.SetLogout);
  const username = useUserStore((state) => state.username);

  return [
    {
      key: 'mainmenu:1',
      label: isPassed ? 'Apps' : '',
      title: 'Apps',
      icon: (
        <AppstoreOutlined
          style={
            isPassed
              ? {}
              : { color: token.colorWhite, position: 'relative', left: '3px', bottom: '-5px', fontSize: '1.15rem' }
          }
        />
      ),
      children: [
        {
          key: 'mainmenu/dashboard:3',
          label: <Link to="/dashboard/auth">Authentication & Authorization</Link>,
        },
        {
          key: 'mainmenu/dashboard:2',
          label: <Link to="/dashboard">Change Log</Link>,
        },
        {
          key: 'mainmenu/dashboard:4',
          label: <Link to="/dashboard/job-scheduler">Job Scheduler</Link>,
        },
      ],
    },
    {
      label: isPassed ? 'Theme' : '',
      key: 'mainmenu:2',
      icon:
        theme.theme === 'dark' ? (
          <MoonOutlined
            style={
              isPassed
                ? {}
                : { color: token.colorWhite, position: 'relative', left: '3px', bottom: '-5px', fontSize: '1.15rem' }
            }
          />
        ) : (
          <SunOutlined
            style={
              isPassed
                ? {}
                : { color: token.colorWhite, position: 'relative', left: '3px', bottom: '-5px', fontSize: '1.15rem' }
            }
          />
        ),
      onClick: theme.theme === 'dark' ? () => theme.SetTheme('light') : () => theme.SetTheme('dark'),
    },
    {
      key: 'mainmenu:3',
      label: isPassed ? 'User' : '',
      title: '',
      icon: (
        <UserOutlined
          style={
            isPassed
              ? {}
              : { color: token.colorWhite, position: 'relative', left: '3px', bottom: '-5px', fontSize: '1.15rem' }
          }
        />
      ),
      children: [
        {
          style: { textAlign: 'center' },
          key: 'mainmenu/logout:1',
          label: (
            <Link to="/dashboard/profile">
              <UserOutlined /> {username.slice(0, 1).toUpperCase() + username.slice(1, username.length).toLowerCase()}
            </Link>
          ),
        },
        {
          key: 'mainmenu/logout:2',
          label: (
            <Button
              size={'middle'}
              block
              type="link"
              icon={<LogoutOutlined />}
              onClick={async () => {
                await mutationLogout
                  .mutateAsync()
                  .then(() => {
                    logout();
                    navigate('/', { replace: true });
                  })
                  .catch(genericServerErrorHandler);
              }}
            >
              Log Out
            </Button>
          ),
        },
      ],
    },
  ];
};
