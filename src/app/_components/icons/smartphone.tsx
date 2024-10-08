import Image from "next/image";

export const Smartphone = () => {
  return (
    <div className="flex justify-center items-center h-auto m-auto p-52">
      <Image
        src="/capture.png"
        width={800}
        height={800}
        alt="Picture of the author"
      />
    </div>
  );
};
