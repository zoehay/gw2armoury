import navbar from "./navbar.module.css";

export const DesktopNav = () => {
  return (
    <div className={navbar.desktopContent}>
      <p>Logo</p>
      <div className={navbar.desktopLinks}>
        <ul className={navbar.link}>Manage Keys</ul>
        <ul className={navbar.link}>Inventory</ul>
      </div>
    </div>
  );
};
