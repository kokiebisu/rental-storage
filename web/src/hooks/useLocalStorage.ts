import { useEffect, useState } from "react";

export const useLocalStorage = () => {
  const [value, setValue] = useState<string | null>(null);

  useEffect(() => {
    const value = localStorage.getItem("authorizationToken");
    setValue(value);
  }, []);

  const setItem = (key: string, value: string) => {
    if (typeof window !== "undefined") {
      localStorage.setItem(key, value);
      setValue(value);
    }
  };

  const getItem = (key: string) => {
    if (typeof window !== "undefined") {
      const value = localStorage.getItem(key);
      setValue(value);
      return value;
    }
  };

  const removeItem = (key: string) => {
    if (typeof window !== "undefined") {
      localStorage.removeItem(key);
      setValue(null);
    }
  };

  return { value, setItem, getItem, removeItem };
};
