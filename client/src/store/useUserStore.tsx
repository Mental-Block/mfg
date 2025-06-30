import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';

import type { User } from 'src/domain/user';

export interface UserStore extends User {
  loggedIn: boolean;
  SetLogout: () => void;
  SetLogedIn: () => void;
  SetState: (state: UserState) => void;
}

type UserState = Omit<UserStore, 'SetState' | 'SetLogout' | 'SetLogedIn'>;

export const defualtUserState: UserState = {
  id: 0,
  username: 'dsa',
  loggedIn: true,
};

export const useUserStore = create<UserStore>()(
  devtools(
    persist(
      (set) => ({
        ...defualtUserState,
        SetState: (state: UserState) => set(state),
        SetLogout: () => set(defualtUserState),
        SetLogedIn: () => set((state) => ({ ...state, loggedIn: true })),
      }),
      {
        name: 'user-storage',
      }
    )
  )
);
