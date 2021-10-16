import Link from "next/link";

export default function Land(): JSX.Element {
  return (
    <div>
      <h1>API Homepage</h1>
      <p>
        This is <code>where you'll see the API</code>.
      </p>
      <p>
        Check out <Link href="/api/spec">bar</Link>.
      </p>
    </div>
  );
}