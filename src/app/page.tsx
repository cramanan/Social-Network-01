import Actualite from "@/components/Actualite";

// import Header from "@/components/Header";

import HomeProfileLayout from "@/layouts/HomeProfileLayout";
// import ChatBox from "@/components/ChatBox";
// import Comment from "@/components/Comment";
// import CreatePost from "@/components/CreatePost";
// import FriendInvite from "@/components/FriendInvite";
// import { ProfileCircle } from "@/components/icons/ProfileCircle";
// import Media from "@/components/Media";
// import Post from "@/components/Post";
// import ProfileBanner from "@/components/ProfileBanner";
// import ProfilePost from "@/components/ProfilePost";
// import ProfileStats from "@/components/ProfileStats";
// import Publication from "@/components/Publication";
// import UserList from "@/components/UserList";
// import Users from "@/components/Users";

export default function Home() {
  return (
    <HomeProfileLayout>
      <main>
        <div className="flex justify-center w-full h-full">
          <Actualite />
        </div>
      </main>
      {/* <main className="flex flex-col gap-5 justify-center items-center">
            <Actualite />
            <Users />
            <Chat />
            <UserList />
            <ChatBox />
            <Comment />
            <CreatePost />
            <Media />
            <Post />
            <ProfileBanner />
            <ProfileStats />
            <ProfilePost />
            <Publication />
            <SideNavMenu />
            <FriendInvite />
            <FriendInviteList />
        </main> */}
    </HomeProfileLayout>
  );
}
