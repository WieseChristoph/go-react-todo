import { type FC } from "react";
import useUser from "~/hooks/useUser";

import Avatar from "../Utils/Avatar";

const ProfileCard: FC = () => {
  const { user } = useUser();

  return (
    <div className="flex flex-col items-center rounded-md border border-gray-700 bg-gray-800 p-6 shadow">
      <Avatar
        className="rounded-full"
        userId={user?.id}
        avatar={user?.avatar}
        size={128}
      />
      <hr className="my-2 w-full border-gray-700" />
      <span className="text-3xl text-white">{user?.global_name}</span>
      <span className="text-xl text-gray-300">{user?.username}</span>
      <span className="text-xl text-gray-300">{user?.email}</span>
      <hr className="my-2 w-full border-gray-700" />
      <div className=" flex flex-col items-center text-sm">
        <span className="text-gray-300">Joined on </span>
        <span className="text-gray-300">
          {user?.created_at ? new Date(user?.created_at).toUTCString() : "-"}
        </span>
      </div>
    </div>
  );
};

export default ProfileCard;
