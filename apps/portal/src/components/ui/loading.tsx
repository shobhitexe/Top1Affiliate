import Image from "next/image";

export default function LoadingSpinner() {
  return (
    <div className="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-1/2">
      <div className="relative">
        <Image
          src={"/images/loader2.svg"}
          alt={"spinner"}
          width={157}
          height={157}
          className="animate-spin"
        />

        <Image
          src={"/images/loader1.svg"}
          alt={"logo"}
          width={86}
          height={86}
          className="absolute left-1/2 -translate-x-1/2 top-1/2 -translate-y-1/2"
        />
      </div>
    </div>
  );
}
