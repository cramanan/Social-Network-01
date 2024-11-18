import Actualite from "@/components/Actualite";
// import Media from "@/components/Media";

import HomeProfileLayout from "@/layouts/HomeProfileLayout";

export default function Home() {
    return (
        <>
            <HomeProfileLayout>
                <Actualite />
            </HomeProfileLayout>
        </>
    );
}
