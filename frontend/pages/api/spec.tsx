import Link from "next/link";

export default function Spec(): JSX.Element {
  return (
    <div>
      <h1>API Spec</h1>
      <p>
        This is <code>pages/foo/bar.tsx</code>.
      </p>
      <p>
        Check out <Link href="/">the homepage</Link>.
      </p>
    </div>
  );
}