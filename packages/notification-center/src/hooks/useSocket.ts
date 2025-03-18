import { useKhulnasoftContext } from './useKhulnasoftContext';

export function useSocket() {
  const { socket } = useKhulnasoftContext();

  return {
    socket,
  };
}
