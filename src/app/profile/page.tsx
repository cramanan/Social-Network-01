import Actualite from "@/components/Actualite";
import ProfileBanner from "@/components/ProfileBanner";
import ProfileStats from "@/components/ProfileStats";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { useAuth } from "@/providers/AuthContext";
import { redirect } from "next/navigation";

export default function Profile() {
  return (
    <>
      <HomeProfileLayout>
        <ProfileBanner />
        <ProfileStats />
        <div>
          <Actualite />
        </div>
      </HomeProfileLayout>
    </>
  );
}
