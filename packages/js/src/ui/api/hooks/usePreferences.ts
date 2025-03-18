import { createEffect, createResource, createSignal, onCleanup, onMount } from 'solid-js';
import { Preference } from '../../../preferences/preference';
import { FetchPreferencesArgs } from '../../../preferences/types';
import { useKhulnasoft } from '../../context';

export const usePreferences = (options?: FetchPreferencesArgs) => {
  const khulnasoft = useKhulnasoft();

  const [loading, setLoading] = createSignal(true);
  const [preferences, { mutate, refetch }] = createResource(options || {}, async ({ tags }) => {
    try {
      const response = await khulnasoft.preferences.list({ tags });

      return response.data;
    } catch (error) {
      console.error('Error fetching preferences:', error);
      throw error;
    }
  });

  onMount(() => {
    const listener = ({ data }: { data: Preference[] }) => {
      if (!data) {
        return;
      }

      mutate(data);
    };

    const cleanup = khulnasoft.on('preferences.list.updated', listener);

    onCleanup(() => cleanup());
  });

  createEffect(() => {
    setLoading(preferences.loading);
  });

  return { preferences, loading, mutate, refetch };
};
