package examples

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Pagination() {
	moreCmd := exec.Command("cmd", "/C", "more")
	moreCmd.Stdin = strings.NewReader(blob)
	moreCmd.Stdout = os.Stdout
	moreCmd.Stderr = os.Stderr
	err := moreCmd.Run()
	if err != nil {
		panic(err)
	}
}

var (
	blob = `Complex Variables
	by Robert B. Ash and W.P. Novinger
	Table Of Contents Chapter 1: Introduction
	1.1 Basic Definitions
	1.2 Further Topology of the Plane
	1.3 Analytic Functions
	1.4 Real-Differentiability and the Cauchy-Riemann Equations 1.5 The Exponential Function
	1.6 Harmonic Functions
	Chapter 2: The Elementary Theory
	2.1 Integration on Paths
	2.2 Power Series
	2.3 The Exponential Function and the Complex Trigonometric Functions 2.4 Further Applications
	Chapter 3: The General Cauchy Theorem
	3.1 Logarithms and Arguments
	3.2 The Index of a Point with Respect to a Closed Curve 3.3 Cauchy’s Theorem
	3.4 Another Version of Cauchy’s Theorem
	Chapter 4: Applications of the Cauchy Theory
	4.1 Singularities
	4.2 Residue Theory
	4.3 The Open mapping Theorem for Analytic Functions 4.4 Linear Fractional Transformations
	4.5 Conformal Mapping
	4.6 Analytic Mappings of One Disk to Another
	1
	2
	4.7 The Poisson Integral formula and its Applications 4.8 The Jensen and Poisson-Jensen Formulas
	4.9 Analytic Continuation
	Chapter 5: Families of Analytic Functions
	5.1 The Spaces A(Ω) and C(Ω)
	5.2 The Riemann Mapping Theorem
	5.3 Extending Conformal Maps to the Boundary
	Chapter 6: Factorization of Analytic Functions
	6.1 Infinite Products
	6.2 Weierstrass Products
	6.3 Mittag-Leffler’s Theorem and Applications
	Chapter 7: The Prime Number Theorem
	7.1 The Riemann Zeta Function
	7.2 An Equivalent Version of the Prime Number Theorem 7.3 Proof of the Prime Number Theorem
	Chapter 1
	Introduction
	The reader is assumed to be familiar with the complex plane C to the extent found in most college algebra texts, and to have had the equivalent of a standard introductory course in real analysis (advanced calculus). Such a course normally includes a discussion of continuity, differentiation, and Riemann-Stieltjes integration of functions from the real line to itself. In addition, there is usually an introductory study of metric spaces and the associated ideas of open and closed sets, connectedness, convergence, compactness, and continuity of functions from one metric space to another. For the purpose of review and to establish notation, some of these concepts are discussed in the following sections.
	1.1 Basic Definitions
	The complex plane C is the set of all ordered pairs (a, b) of real numbers, with addition and multiplication defined by
	(a,b)+(c,d) = (a+c,b+d) and (a,b)(c,d) = (ac−bd,ad+bc).
	If i = (0,1) and the real number a is identified with (a,0), then (a,b) = a + bi. The expression a + bi can be manipulated as if it were an ordinary binomial expression of real numbers, subject to the relation i2 = −1. With the above definitions of addition and multiplication, C is a field.
	If z = a+bi, then a is called the real part of z, written a = Rez, and b is called the imaginary part of z, written b = Im z. The absolute value or magnitude or modulus of z is defined as (a2 + b2)1/2. A complex number with magnitude 1 is said to be unimodular. An argument of z (written arg z) is defined as the angle which the line segment from (0, 0) to (a, b) makes with the positive real axis. The argument is not unique, but is determined up to a multiple of 2π.
	If r is the magnitude of z and θ is an argument of z, we may write z = r(cos θ + i sin θ)
	and it follows from trigonometric identities that
	|z1z2| = |z1||z2| and
	arg(z1z2) = arg z1 + arg z2
	1
	2 CHAPTER 1. INTRODUCTION
	(that is, if θk is an argument of zk,k = 1,2, then θ1 + θ2 is an argument of z1z2). If z2 ̸= 0, then arg(z1/z2) = arg(z1) − arg(z2). If z = a + bi, the conjugate of z is defined as z = a − bi, and we have the following properties:
	|z|=|z|, argz=−argz, z1+z2 =z1+z2, z1−z2 =z1−z2, z1z2 =z1z2, Rez=(z+z)/2, Imz=(z−z)/2i, zz=|z|2.
	The distance between two complex numbers z1 and z2 is defined as d(z1, z2) = |z1 − z2|. So d(z1,z2) is simply the Euclidean distance between z1 and z2 regarded as points in the plane. Thus d defines a metric on C, and furthermore, d is complete, that is, every Cauchy sequence converges. If z1, z2, . . . is sequence of complex numbers, then zn → z if and only if Re zn → Re z and Im zn → Im z. We say that zn → ∞ if the sequence of real numbers |zn| approaches +∞.
	Many of the above results are illustrated in the following analytical proof of the triangle inequality:
	|z1 +z2|≤|z1|+|z2| for all z1,z2 ∈C.
	The geometric interpretation is that the length of a side of a triangle cannot exceed the sum of the lengths of the other two sides. See Figure 1.1.1, which illustrates the familiar representation of complex numbers as vectors in the plane.
	The proof is as follows:
	oo􏱷 z1+z2 oo 􏱷
	ooo 􏱷􏱷 ooooo // 􏱷􏱷􏱷 z2
	z1
	Figure 1.1.1
	|z1 +z2|2 =(z1 +z2)(z1 +z2)=|z1|2 +|z2|2 +z1z2 +z1z2
	= |z1|2 + |z2|2 + z1z2 + z1z2 = |z1|2 + |z2|2 + 2 Re(z1z2) ≤ |z1|2 + |z2|2 + 2|z1z2| = (|z1| + |z2|)2.
	The proof is completed by taking the square root of both sides.
	If a and b are complex numbers, [a, b] denotes the closed line segment with endpoints
	a and b. If t1 and t2 are arbitrary real numbers with t1 < t2, then we may write [a,b]={a+ t−t1 (b−a):t1 ≤t≤t2}.
	t2 − t1
	The notation is extended as follows. If a1, a2, . . . , an+1 are points in C, a polygon from
	a1 to an+1 (or a polygon joining a1 to an+1) is defined as 􏱽n
	often abbreviated as [a1, . . . , an+1].
	[aj,aj+1], j=1
	77 o??
	
	1.2. FURTHER TOPOLOGY OF THE PLANE 3 1.2 Further Topology of the Plane
	Recall that two subsets S1 and S2 of a metric space are separated if there are open sets G1 ⊇S1 andG2 ⊇S2 suchthatG1∩S2 =G2∩S1 =∅,theemptyset. Asetis connected iff it cannot be written as the union of two nonempty separated sets. An open (respectively closed) set is connected iff it is not the union of two nonempty disjoint open (respectively closed) sets.
	1.2.1 Definition
	A set S ⊆ C is said to be polygonally connected if each pair of points in S can be joined by a polygon that lies in S.
	Polygonal connectedness is a special case of path (or arcwise) connectedness, and it follows that a polygonally connected set, in particular a polygon itself, is connected. We will prove in Theorem 1.2.3 that any open connected set is polygonally connected.
	1.2.2 Definitions
	If a ∈ C and r > 0, then D(a,r) is the open disk with center a and radius r; thus D(a,r)={z:|z−a|<r}. Thecloseddisk{z:|z−a|≤r}isdenotedbyD(a,r),and C(a,r) is the circle with center a and radius r.
	1.2.3 Theorem
	If Ω is an open subset of C, then Ω is connected iff Ω is polygonally connected.
	Proof. IfΩisconnectedanda∈Ω,letΩ1 bethesetofallzinΩsuchthatthereisa polygoninΩfromatoz,andletΩ2 =Ω\Ω1. Ifz∈Ω1,thereisanopendiskD(z,r)⊆Ω (because Ω is open). If w ∈ D(z,r), a polygon from a to z can be extended to w, and it follows that D(z,r) ⊆ Ω1, proving that Ω1 is open. Similarly, Ω2 is open. (Suppose z ∈ Ω2, and choose D(z, r) ⊆ Ω. Then D(z, r) ⊆ Ω2 as before.)
	Thus Ω1 and Ω2 are disjoint open sets, and Ω1 ̸= ∅ because a ∈ Ω1. Since Ω is connected we must have Ω2 = ∅, so that Ω1 = Ω. Therefore Ω is polygonally connected. The converse assertion follows because any polygonally connected set is connected. ♣
	1.2.4Definitions
	A region in C is an open connected subset of C. A set E ⊆ C is convex if for each pair of points a,b ∈ E, we have [a,b] ⊆ E; E is starlike if there is a point a ∈ E (called a star center) such that [a,z] ⊆ E for each z ∈ E. Note that any nonempty convex set is starlike and that starlike sets are polygonally connected.
	
	4 CHAPTER 1. INTRODUCTION 1.3 Analytic Functions
	1.3.1 Definition
	Let f : Ω → C, where Ω is a subset of C. We say that f is complex-differentiable at the pointz0 ∈Ωifforsomeλ∈Cwehave
	 or equivalently,
	lim f(z0 +h)−f(z0) = λ (1) h→0 h
	lim f(z)−f(z0)=λ. (2) z→z0 z − z0
	 Conditions (3), (4) and (5) below are also equivalent to (1), and are sometimes easier to apply.
	lim f(z0 +hn)−f(z0) =λ (3) n→∞ hn
	for each sequence {hn} such that z0 + hn ∈ Ω \ {z0} and hn → 0 as n → ∞.
	lim f(zn)−f(z0)=λ (4)
	n→∞ zn − z0 foreachsequence{zn}suchthatzn ∈Ω\{z0}andzn →z0 asn→∞.
	f(z) = f(z0) + (z − z0)(λ + ε(z)) (5) forallz∈Ω,whereε:Ω→Ciscontinuousatz0 andε(z0)=0.
	To show that (1) and (5) are equivalent, just note that ε may be written in terms of f as follows:
	􏱾f(z)−f(z0)−λ ifz̸=z0
	at z0.
	If f is complex-differentiable at every point of Ω, f is said to be analytic or holomorphic
	on Ω. Analytic functions are the basic objects of study in complex variables.
	Analyticity on a nonopen set S ⊆ C means analyticity on an open set Ω ⊇ S. In particular, f is analytic at a point z0 iff f is analytic on an open set Ω with z0 ∈ Ω. If f1 and f2 are analytic on Ω, so are f1 + f2, f1 − f2, kf1 for k ∈ C, f1f2, and f1/f2 (provided that f2 is never 0 on Ω). Furthermore,
	(f1 +f2)′ = f1′ +f2′, (f1 −f2)′ = f1′ −f2′, (kf1)′ = kf1′
	   ε(z) = z−z0 0
	if z = z0.
	The number λ is unique. It is usually written as f′(z0), and is called the derivative of f
	􏱿f 􏲀′
	(f1f2)′ =f1f2′ +f1′f2, 1 f2
	f f′ −f f′ = 2 1 1 2.
	f2
	 
	1.4. REAL-DIFFERENTIABILITY AND THE CAUCHY-RIEMANN EQUATIONS 5 The proofs are identical to the corresponding proofs for functions from R to R.
	Since d (z) = 1 by direct computation, we may use the rule for differentiating a dz
	product (just as in the real case) to obtain
	d (zn) = nzn−1, n = 0,1,...
	dz
	This extends to n = −1, −2, . . . using the quotient rule.
	If f is analytic on Ω and g is analytic on f(Ω) = {f(z) : z ∈ Ω}, then the composition g◦f is analytic on Ω and
	d g(f(z)) = g′(f(z)f′(z) dz
	just as in the real variable case.
	As an example of the use of Condition (4) of (1.3.1), we now prove a result that will be useful later in studying certain inverse functions.
	1.3.2 Theorem
	Let g be analytic on the open set Ω1, and let f be a continuous complex-valued function on the open set Ω. Assume
	(i) f(Ω) ⊆ Ω1,
	(ii) g′ is never 0,
	(iii) g(f(z)) = z for all z ∈ Ω (thus f is 1-1).
	Then f is analytic on Ω and f′ = 1/(g′ ◦ f).
	Proof. Letz0 ∈Ω,andlet{zn}beasequenceinΩ\{z0}withzn →z0. Then
	f(zn) − f(z0) f(zn) − f(z0) 􏲁g(f(zn)) − g(f(z0))􏲂−1 z −z = g(f(z ))−g(f(z )) = f(z )−f(z ) .
	   n0n0n0
	(Note that f(zn) ̸= f(z0) since f is 1-1 and zn ̸= z0.) By continuity of f at z0, the expression in brackets approaches g′(f(z0)) as n → ∞. Since g′(f(z0)) ̸= 0, the result follows. ♣
	1.4Real-DifferentiabilityandtheCauchy-RiemannEqua- tions
	Let f : Ω → C, and set u = Ref,v = Imf. Then u and v are real-valued functions on Ω and f = u + iv. In this section we are interested in the relation between f and its real and imaginary parts u and v. For example, f is continuous at a point z0 iff both u and v are continuous at z0. Relations involving derivatives will be more significant for us, and for this it is convenient to be able to express the idea of differentiability of real-valued function of two real variables by means of a single formula, without having to consider partial derivatives separately. We do this by means of a condition analogous to (5) of (1.3.1).
	
	6 CHAPTER 1. INTRODUCTION Convention
	From now on, Ω will denote an open subset of C, unless otherwise specified. 1.4.1 Definition
	Letg:Ω→R. Wesaythatgisreal-differentiableatz0 =x0+iy0 ∈Ωifthereexist real numbers A and B, and real functions ε1 and ε2 defined on a neighborhood of (x0, y0), such that ε1 and ε2 are continuous at (x0, y0), ε1(x0, y0) = ε2(x0, y0) = 0, and
	g(x, y) = g(x0, y0) + (x − x0)[A + ε1(x, y)] + (y − y0)[B + ε2(x, y)] for all (x, y) in the above neighborhood of (x0, y0).
	It follows from the definition that if g is real-differentiable at (x0 , y0 ), then the partial derivatives of g exist at (x0,y0) and
	∂g(x0,y0) = A, ∂g(x0,y0) = B. ∂x ∂y
	If, on the other hand, ∂g and ∂g exist at (x0, y0) and one of these exists in a neighborhood ∂x ∂y
	of (x0, y0) and is continuous at (x0, y0), then g is real-differentiable at (x0, y0). To verify this, assume that ∂ g is continuous at (x0 , y0 ), and write
	∂x
	g(x,y)−g(x0,y0) = g(x,y)−g(x0,y)+g(x0,y)−g(x0,y0).
	Now apply the mean value theorem and the definition of partial derivative respectively
	(Problem 4).
	1.4.2 Theorem
	Let f : Ω → C,u = Ref,v = Imf. Then f is complex-differentiable at (x0,y0) iff u and v are real-differentiable at (x0,y0) and the Cauchy-Riemann equations
	∂u = ∂v, ∂v =−∂u ∂x ∂y ∂x ∂y
	are satisfied at (x0, y0). Furthermore, if z0 = x0 + iy0, we have
	f′(z0) = ∂u(x0,y0)+i∂v(x0,y0) = ∂v(x0,y0)−i∂u(x0,y0).
	Proof. Assume f complex-differentiable at z0, and let ε be the function supplied by (5) of (1.3.1). Define ε1(x, , y) = Re ε(x, y), ε2(x, y) = Im ε(x, y). If we take real parts of both sides of the equation
	we obtain
	f(x) = f(z0) + (z − z0)(f′(z0) + ε(z)) (1)
	u(x, y) = u(x0, y0) + (x − x0)[Re f′(z0) + ε1(x, y)]
	+ (y − y0)[− Im f′(z0) − ε2(x, y)].
	∂x ∂x ∂y ∂y
	
	1.5. THE EXPONENTIAL FUNCTION 7 It follows that u is real-differentiable at (x0,y0) with
	∂u(x0,y0) = Ref′(z0), ∂u(x0,y0) = −Imf′(z0). (2) ∂x ∂y
	Similarly, take imaginary parts of both sides of (1) to obtain
	v(x, y) = v(x0, y0) + (x − x0)[Im f′(z0) + ε2(x, y)]
	and conclude that
	+ (y − y0)[Re f′(z0) + ε1(x, y)]
	∂v (x0, y0) = Im f′(z0), ∂v (x0, y0) = Re f′(z0). (3)
	∂x ∂y
	The Cauchy-Riemann equations and the desired formulas for f′(z0) follow from (2) and
	(3).
	Conversely, suppose that u and v are real-differentiable at (x0,y0) and satisfy the Cauchy-Riemann equations there. Then we may write equations of the form
	u(x,y) = u(x0,y0)+(x−x0)[∂u(x0,y0)+ε1(x,y)] ∂x
	+ (y − y0)[ ∂u (x0, y0) + ε2(x, y)], (4) ∂y
	v(x,y) = v(x0,y0)+(x−x0)[∂v(x0,y0)+ε3(x,y)] ∂x
	+ (y − y0)[ ∂v (x0, y0) + ε4(x, y)]. (5) ∂y
	Since f = u + iv, (4) and (5) along with the Cauchy-Riemann equations yield f (z ) = f (z0 ) + (z − z0 )[ ∂ u (x0 , y0 ) + i ∂ v (x0 , y0 ) + ε(z )]
	where, at least in a neighborhood of z0,
	􏲁x−x 􏲂 ε(z) = 0
	z − z0
	􏲁y−y 􏲂
	0 [ε2(x,y) + iε4(x,y)] if z ̸= z0; ε(z0) = 0.
	[ε1(x,y) + iε3(x,y)] +
	It follows that f is complex-differentiable at z0. ♣
	∂x ∂x
	z − z0
	1.5 The Exponential Function
	In this section we extend the domain of definition of the exponential function (as normally encountered in calculus) from the real line to the entire complex plane. If we require that the basic rules for manipulating exponentials carry over to the extended function, there is
	
	8 CHAPTER 1. INTRODUCTION only one possible way to define exp(z) for z = x+iy ∈ C. Consider the following sequence
	of “equations” that exp should satisfy:
	exp(z) = exp(x + iy)
	“ = ” exp(x) exp(iy)
	􏱿 (iy)2 􏲀 “=”ex 1+iy+ 2! +···
	􏲁􏱿 y2 y4 􏲀 􏱿 y3
	“ = ” ex 1 − 2! + 4! − · · · + i y − 3! + 5! − · · ·
	“ = ” ex(cos y + i sin y).
	Thus we have only one candidate for the role of exp on C.
	1.5.1 Definition
	If z = x+iy ∈ C, let exp(z) = ex(cosy+isiny). Note that if z = x ∈ R, then exp(z) = ex so exp is indeed a extension of the real exponential function.
	1.5.2 Theorem
	The exponential function is analytic on C and d exp(z) = exp(z) for all z. dz
	Proof. The real and imaginary parts of exp(x + iy) are, respectively, u(x, y) = ex cos y and v(x, y) = ex sin y. At any point (x0, y0), u and v are real-differentiable (see Problem 4) and satisfy the Cauchy-Riemann equations there. The result follows from (1.4.2). ♣
	Functions such as exp and the polynomials that are analytic on C are called entire functions.
	The exponential function is of fundamental importance in mathematics, and the in- vestigation of its properties will be continued in Section 2.3.
	1.6 Harmonic Functions 1.6.1 Definition
	A function g : Ω → R is said to be harmonic on Ω if g has continuous first and second order partial derivatives on Ω and satisfies Laplace’s equation
	∂2g + ∂2g = 0 ∂x2 ∂y2
	on all of Ω.
	After some additional properties of analytic functions have been developed, we will be able to prove that the real and imaginary parts of an analytic function on Ω are harmonic on Ω. The following theorem is a partial converse to that result, namely that a harmonic on Ω is locally the real part of an analytic function.
	y5 􏲀􏲂
	
	1.6. HARMONIC FUNCTIONS 9 1.6.2 Theorem
	Suppose u : Ω → R is harmonic on Ω, and D is any open disk contained in Ω. Then there exists a function v : D → R such that u + iv is analytic on D.
	The function v is called a harmonic conjugate of u.
	Proof. Consider the differential Pdx + Qdy, where P = −∂u, Q = ∂u. Since u is
	∂y
	harmonic, P and Q have continuous partial derivatives on Ω and ∂ P ∂y
	∂x
	=
	(from calculus) that P dx + Qdy is a locally exact differential. In other words, there is a
	functionv:D→Rsuchthatdv=Pdx+Qdy. ButthisjustmeansthatonDwehave ∂v =P =−∂u and ∂v =Q= ∂u.
	∂x ∂y ∂y ∂x Hence by (1.4.2) (and Problem 4), u + iv is analytic on D.
	Problems
	1. Prove the parallelogram law |z1 + z2|2 + |z1 − z2|2 = 2[|z1|2 + |z2|2] and give a geometric interpretation.
	2. Showthat|z1+z2|=|z1|+|z2|iffz1 andz2 lieonacommonrayfrom0iffoneof z1 or z2 is a nonnegative multiple of the other.
	3. Let z1 and z2 be nonzero complex numbers, and let θ,0 ≤ θ ≤ π, be the angle between them. Show that
	(a) Re z1z2 = |z1||z2| cos θ, Im z1z2 = ±|z1||z2| sin θ, and consequently
	(b) The area of the triangle formed by z1, z2 and z2 − z1 is | Im z1z2|/2.
	4. Letg:Ω→Rbesuchthat ∂g and ∂g existat(x0,y0)∈Ω,andsupposethatone
	∂x ∂y
	of these partials exists in a neighborhood of (x0,y0) and is continuous at (x0,y0). Show that g is real-differentiable at (x0,y0).
	5. Let f(x) = z,z ∈ C. Show that although f is continuous everywhere, it is nowhere differentiable.
	6. Let f(z) = |z|2,z ∈ C. Show that f is complex-differentiable at z = 0, but nowhere else. 􏲃 ∂u ∂u
	7. Let u(x,y) = |xy|,(x,y) ∈ C. Show that ∂x and ∂y both exist at (0,0), but u is not real-differentiable at (0,0).
	8. Show that the field of complex numbers is isomorphic to the set of matrices of the
	form
	with a, b ∈ R.
	􏲁􏲂
	ab −b a
	9. Show that the complex field cannot be ordered. That is, there is no subset P ⊆ C of “positive elements” such that
	(a) P is closed under addition and multiplication. (b)Ifz∈P,thenexactlyoneoftherelationsz∈P, z=0, −z∈P holds.
	∂ Q . It follows ∂x
	
	10 CHAPTER 1. INTRODUCTION
	10. (A characterization of absolute value) Show that there is a unique function α : C → R such that
	(i) α(x) = x for all real x ≥ 0;
	(ii) α(zw) = α(z)α(w), z, w ∈ C;
	(iii) α is bounded on the unit circle C(0,1).
	Hint: First show that α(z) = 1 for |z| = 1.
	11. (Another characterization of absolute value) Show that there is a unique function α : C → R such that
	(i) α(x) = x for all real x ≥ 0;
	(ii) α(zw) = α(z)α(w), z, w ∈ C;
	(iii) α(z + w) ≤ α(z) + α(w),z,w ∈ C.
	12. Let α be a complex number with |α| < 1. Prove that
	􏲄􏲄􏲄z−α􏲄􏲄􏲄=1 iff |z|=1. 􏲄1−αz􏲄
	13. Supposez∈C,z̸=0. Showthatz+z1 isrealiffImz=0or|z|=1.
	14. In each case show that u is harmonic and find the harmonic conjugate v such that v(0, 0) = 0.
	(i) u(x, y) = ey cos x;
	(ii) u(x, y) = 2x − x3 + 3xy2.
	15. Leta,b∈Cwitha̸=0,andletT(z)=az+b,z∈C.
	(i) Show that T maps the circle C(z0, r) onto the circle C(T (z0), r|a|). (ii) For which choices of a and b will T map C(0, 1) onto C(1 + i, 2)? (iii) In (ii), is it possible to choose a and b so that T(1) = −1 + 3i?
	16. Show that f (z) = eRe z is nowhere complex-differentiable.
	17. Let f be a complex-valued function defined on an open set Ω that is symmetric with respect to the real line, that is, z ∈ Ω implies z ∈ Ω. (Examples are C and D(x, r) where x ∈ R.) Set g(z) = f(z), and show that g is analytic on Ω if and only if f is analytic on Ω.
	18. Show that an equation for the circle C(z0, r) is zz − z0z − z0z + z0z0 = r2.
	19. (Enestrom’s theorem) Suppose that P (z) = a0 + a1z + · · · + anzn, where n ≥ 1 and a0 ≥a1 ≥a2 ≥···≥an >0. ProvethatthezerosofthepolynomialP(z)alllie outside the open unit disk D(0, 1).
	Hint: Look at (1 − z)P(z), and show that (1 − z)P(z) = 0 implies that a0 = (a0 − a1)z + (a1 − a2)z2 + · · · + (an−1 − an)zn + anzn+1, which is impossible for |z| < 1.
	20. Continuing Problem 19, show that if aj−1 > aj for all j, then all the zeros of P(z) must be outside the closed unit disk D(0, 1).
	Hint: If the last equation of Problem 19 holds for some z with |z| ≤ 1, then z = 1.`
)
