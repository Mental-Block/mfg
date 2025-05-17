import { create } from 'zustand';
import { devtools, persist } from 'zustand/middleware';

export type Theme = 'light' | 'dark';

export interface ThemeStore {
  theme: Theme;
  SetTheme: (state: Theme) => void;
}

export const useThemeStore = create<ThemeStore>()(
  devtools(
    persist(
      (set) => ({
        theme: 'light',
        SetTheme: (theme: Theme) => set({ theme }),
      }),
      {
        name: 'user-storage',
      }
    )
  )
);
