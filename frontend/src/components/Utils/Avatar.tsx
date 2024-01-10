import { type FC } from "react";

type Props = {
  userId?: number;
  avatar?: string;
  size?: number;
  className?: string;
};

const Avatar: FC<Props> = ({ userId, avatar, size = 48, className = "" }) => {
  return (
    <img
      className={className}
      src={
        userId && avatar
          ? `https://cdn.discordapp.com/avatars/${userId}/${avatar}.png`
          : "https://cdn.discordapp.com/embed/avatars/1.png"
      }
      alt="Avatar"
      width={size}
      height={size}
    />
  );
};

export default Avatar;
