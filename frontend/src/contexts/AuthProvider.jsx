import React, { useState, useEffect, useCallback, useRef } from "react";
import { getMe } from "../service/user";
import { AuthContext } from "./authContext";

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const hasFetched = useRef(false);

  const fetchUserOnce = useCallback(async () => {
    if (hasFetched.current) return;
    hasFetched.current = true;

    try {
      const response = await getMe();
      setUser(response.data);
    } catch (error) {
      console.error("Error fetching user:", error);
      setUser(null);
    } finally {
      setLoading(false);
    }
  }, []);

  const refetchUser = useCallback(async () => {
    try {
      const response = await getMe();
      setUser(response.data.data);
    } catch (error) {
      console.error("Error refetching user:", error);
      setUser(null);
    }
  }, []);

  useEffect(() => {
    fetchUserOnce();
  }, [fetchUserOnce]);

  return (
    <AuthContext.Provider value={{ user, loading, refetchUser }}>
      {children}
    </AuthContext.Provider>
  );
};
