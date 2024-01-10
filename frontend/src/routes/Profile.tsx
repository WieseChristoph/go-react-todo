import { type FC } from "react";
import useUser from "~/hooks/useUser";

import ProfileCard from "~/components/Profile/ProfileCard";

const Profile: FC = () => {
  const { user } = useUser();

  if (user)
    return (
      <div className="flex justify-center">
        <ProfileCard />
      </div>
    );
  else
    return (
      <div className="w-full text-center text-xl text-red-500">
        You need to login to see your profile
      </div>
    );
};

export default Profile;
