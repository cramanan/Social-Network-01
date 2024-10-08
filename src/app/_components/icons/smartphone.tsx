import Image from "next/image";

export const Smartphone = () => {
  return (
    <div className="flex align-items-center ">
      <Image
        src="/capture.png"
        width={900}
        height={900}
        alt="Picture of the author"
      />
    </div>
  );
};
