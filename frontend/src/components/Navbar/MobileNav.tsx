import navbar from "./navbar.module.css";
import { useState } from "react";

export const MobileNav = () => {
  let [showMenu, setShowMenu] = useState(false);
  let alt;

  if (!showMenu) {
    alt = "options";
  } else {
    alt = "close";
  }

  const handleClose = () => {
    if (showMenu) {
      setShowMenu(false);
    }
  };

  return (
    <>
      <button onClick={() => setShowMenu(!showMenu)}> </button>
      {showMenu && (
        <div className={navbar.mobileDropdownWrapper} onClick={handleClose}>
          <div className={navbar.mobileDropdown}></div>
        </div>
      )}
    </>
  );
};
