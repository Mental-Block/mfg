import React from 'react';
import { useNavigate, Link, useLocation } from 'react-router';
import { useMutation } from '@tanstack/react-query';
import { Button, MenuProps, theme as antdTheme } from 'antd';
import { ItemType, MenuItemType } from 'antd/es/menu/interface';

import { useWindowResize } from 'src/hooks/useWindowResize';
import { useThemeStore } from 'src/store/useThemeStore';
import { useLogoutMutation } from 'src/features/authentication/hooks/useLogoutMutation';
import { useUserStore, defaultUserState } from 'src/store/useUserStore';
import { genericServerErrorHandler } from 'src/utils/error';

import {
  BuildOutlined,
  NotificationOutlined,
  MinusOutlined,
  PlusOutlined,
  SwapOutlined,
  PieChartOutlined,
  EditOutlined,
  LineOutlined,
  HomeOutlined,
  SettingOutlined,
  MoonOutlined,
  SunOutlined,
  UserOutlined,
  LogoutOutlined,
  AppstoreOutlined,
} from '@ant-design/icons';
type MenuItem = Required<MenuProps>['items'][number];

export const useNavMenu = (): MenuItem[] | undefined => {
  const { token } = antdTheme.useToken();
  const navigate = useNavigate();
  const userStore = useUserStore();
  const theme = useThemeStore();
  const mutationLogout = useLogoutMutation();
  const { username } = useUserStore();
  const [isPassed] = useWindowResize('md', false);

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
          key: 'mainmenu/dashboard:2',
          label: <Link to="/dashboard">Change Log</Link>,
        },
        {
          key: 'mainmenu/dashboard:3',
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
              onClick={async () => {
                await mutationLogout
                  .mutateAsync()
                  .then(() => {
                    userStore.SetState(defaultUserState);
                    navigate('/', { replace: true });
                  })
                  .catch(genericServerErrorHandler);
              }}
              block
              type="link"
              icon={<LogoutOutlined />}
            >
              Log Out
            </Button>
          ),
        },
      ],
    },
    // {
    //   key: 'mainmenu:3',
    //   label: `${Username.slice(0, 1).toUpperCase() + Username.slice(1, Username.length).toLowerCase()}`,
    //   title: 'Apps',
    //   icon: <UserOutlined />,
    //   children: [
    //     {
    //       /*
    //       If the menu changes at all please change the nth-child number in header.css
    //       Its a really crappy solution. I know
    //     */
    //       key: 'apps/logout:1',
    //       style: {
    //         padding: isMedium ? '0px' : '0 8px',
    //       },
    //       label: (
    //         <Button
    //           style={{
    //             marginTop: isMedium ? '2px' : 0,
    //             fontSize: isMedium ? '12px' : 'auto',
    //           }}
    //           size={isMedium ? 'large' : 'middle'}
    //           onClick={async () => {
    //             await mutationLogout
    //               .mutateAsync()
    //               .then(() => {
    //                 userStore.SetState(defaultUserState);
    //                 navigate('/', { replace: true });
    //               })
    //               .catch(genericServerErrorHandler);
    //           }}
    //           block
    //           type="primary"
    //         >
    //           Log Out
    //         </Button>
    //       ),
    //     },
    //   ],
    // },
  ];
};
