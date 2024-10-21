import Actualite from "@/components/Actualite";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";

import HomeProfileLayout from "@/layouts/HomeProfileLayout";

export default function Profile() {
  return (
    <HomeProfileLayout>
      <ProfileBanner />
      <ProfileStats />
      <Actualite />
    </HomeProfileLayout>
  );
}
