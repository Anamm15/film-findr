import { useState, useEffect } from "react";
import { getUserByUsername } from "../../service/user";

export function useFetchUserByUsername(username) {
   const [user, setUser] = useState(null);
   const [loading, setLoading] = useState(false);

   useEffect(() => {
      const fetchUser = async () => {
         if (!username) return;
         setLoading(true);
         try {
            const response = await getUserByUsername(username);
            setUser(response.data);
         } catch (error) {
            console.error("Error fetching user:", error);
         } finally {
            setLoading(false);
         }
      };

      fetchUser();
   }, [username]);

   return { user, loading };
}