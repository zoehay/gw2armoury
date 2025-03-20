import { DesktopNav } from "./DesktopNav";
import { MobileNav } from "./MobileNav";
import navbar from "./navbar.module.css";

export const Navbar = () => {
  return (
    <div className={navbar.nav}>
      <div className={navbar.mobileDiv}>
        <MobileNav></MobileNav>
      </div>
      <div className={navbar.desktopDiv}>
        <DesktopNav></DesktopNav>
      </div>
    </div>
  );
};
