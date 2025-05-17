import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';

import type { User } from 'src/domain/user';

export interface UserState extends User {
  loggedIn: boolean;
}

export const defaultUserState: UserState = {
  id: 0,
  username: '',
  roles: [],
  loggedIn: false,
};

export interface UserStore extends UserState {
  SetState: (state: UserState) => void;
  SetLoggedIn: (state: boolean) => void;
}

export const useUserStore = create<UserStore>()(
  devtools(
    persist(
      (set) => ({
        ...defaultUserState,
        SetState: (state: UserState) => set(state),
        SetLoggedIn: (value: boolean) => set((state) => ({ ...state, LoggedIn: value })),
      }),
      {
        name: 'user-storage',
      }
    )
  )
);
