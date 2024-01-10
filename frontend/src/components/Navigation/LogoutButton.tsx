import type { FC } from "react";

type Props = {
  href?: string;
  onClick?: () => void;
};

const LogoutButton: FC<Props> = ({ href, onClick }) => {
  return (
    <a
      href={href}
      onClick={onClick}
      className=" text-white hover:cursor-pointer"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
        width={32}
        height={32}
      >
        <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
        <polyline points="16 17 21 12 16 7" />
        <line x1="21" x2="9" y1="12" y2="12" />
      </svg>
    </a>
  );
};

export default LogoutButton;
