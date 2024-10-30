import Actualite from "@/components/Actualite";

// import Header from "@/components/Header";

import HomeProfileLayout from "@/layouts/HomeProfileLayout";
// import ChatBox from "@/components/ChatBox";
// import ChatList from "@/components/ChatList";
// import Comment from "@/components/Comment";
// import CreatePost from "@/components/CreatePost";
// import FindUser from "@/components/FindUser";
// import FriendInvite from "@/components/FriendInvite";
// import FriendInviteList from "@/components/FriendInviteList";
// import { ProfileCircle } from "@/components/icons/ProfileCircle";
// import Media from "@/components/Media";
// import Post from "@/components/Post";
// import ProfileBanner from "@/components/ProfileBanner";
// import ProfilePost from "@/components/ProfilePost";
// import ProfileStats from "@/components/ProfileStats";
// import Publication from "@/components/Publication";
// import SideNavMenu from "@/components/SideNavMenu";
// import { useAuth } from "@/providers/AuthContext";

export default function Home() {
  return (
    <>
      <HomeProfileLayout>
        <main>
          <Actualite />
        </main>
      </HomeProfileLayout>
    </>
  );
}
