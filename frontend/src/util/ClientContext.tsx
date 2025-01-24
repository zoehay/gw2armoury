import React, { createContext } from "react";
import { Client } from "./Client";

const mode = import.meta.env.MODE;
const client = new Client(mode);

export const ClientContext = createContext(client);

export const ClientProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  return (
    <ClientContext.Provider value={client}>{children}</ClientContext.Provider>
  );
};
