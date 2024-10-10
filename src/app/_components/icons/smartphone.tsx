import Image from "next/image";

export const Smartphone = () => {
  return (
    <div className="flex items-center justify-center  ">
      <Image
        src="/capture.png"
        width={350}
        height={350}
        alt="Picture of the author"
      />
    </div>
  );
};
