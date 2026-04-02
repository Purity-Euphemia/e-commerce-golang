import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { addReview, addToCart, getProduct, getProductReviews } from '../api';
import { Product, Review } from '../types';
import Loading from '../components/Loading';
import { useAuth } from '../hooks/useAuth';

export default function ProductDetail() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [product, setProduct] = useState<Product | null>(null);
  const [quantity, setQuantity] = useState(1);
  const [reviews, setReviews] = useState<Review[]>([]);
  const [reviewComment, setReviewComment] = useState('');
  const [reviewTitle, setReviewTitle] = useState('');
  const [reviewRating, setReviewRating] = useState(5);
  const [reviewError, setReviewError] = useState('');
  const [loading, setLoading] = useState(true);
  const { token } = useAuth();

  useEffect(() => {
    const load = async () => {
      if (id) {
        try {
          const [data, reviewData] = await Promise.all([getProduct(id), getProductReviews(Number(id))]);
          setProduct(data);
          setReviews(reviewData);
        } catch (error) {
          console.error(error);
        } finally {
          setLoading(false);
        }
      }
    };
    load();
  }, [id]);

  const handleAddToCart = async () => {
    if (!product) return;
    try {
      await addToCart(product.id, quantity);
      navigate('/cart');
    } catch (error) {
      console.error(error);
      alert('Please login or try again.');
    }
  };

  if (loading) return <Loading />;

  if (!product) {
    return <div className="px-4 py-10 text-center text-slate-700">Product not found.</div>;
  }

  return (
    <section className="px-4 py-10 sm:px-6 lg:px-8">
      <div className="mx-auto grid max-w-6xl gap-10 lg:grid-cols-[1.2fr_0.8fr]">
        <div className="rounded-[2rem] bg-white p-8 shadow-soft">
          <img src={product.image || 'https://via.placeholder.com/720x480'} alt={product.name} className="w-full rounded-[1.75rem] object-cover" />
          <div className="mt-8 space-y-4">
            <h1 className="text-3xl font-semibold text-slate-900">{product.name}</h1>
            <p className="text-lg text-slate-600">{product.description || 'A premium product ready for your customers.'}</p>
            <div className="flex flex-wrap items-center gap-3 text-slate-900">
              <span className="text-3xl font-bold">${product.discount_price ?? product.price}</span>
              {product.discount_price ? <span className="text-sm text-slate-500 line-through">${product.price}</span> : null}
            </div>
          </div>
        </div>

        <div className="space-y-6 rounded-[2rem] bg-white p-8 shadow-soft">
          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <p className="text-sm uppercase tracking-[.2em] text-slate-500">Product details</p>
              <span className="text-sm font-semibold text-slate-900">{product.stock ?? 0} in stock</span>
            </div>
            <div className="rounded-3xl border border-slate-200 bg-slate-50 p-4">
              <p className="text-sm text-slate-600">SKU: {product.sku || 'SW-001'}</p>
              <p className="mt-2 text-sm text-slate-600">Category ID: {product.category_id || 'N/A'}</p>
            </div>
          </div>
          <div className="space-y-4">
            <label className="block text-sm font-medium text-slate-700">Quantity</label>
            <input
              type="number"
              value={quantity}
              min={1}
              onChange={(e) => setQuantity(Number(e.target.value))}
              className="w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
            />
          </div>
          <button
            type="button"
            onClick={handleAddToCart}
            className="w-full rounded-3xl bg-slate-900 px-5 py-3 text-sm font-semibold uppercase tracking-[.08em] text-white transition hover:bg-slate-700"
          >
            Add to cart
          </button>
        </div>
      </div>

      <div className="mx-auto max-w-6xl space-y-8 px-4 py-10 sm:px-6 lg:px-8">
        <div className="rounded-[2rem] bg-white p-8 shadow-soft">
          <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
            <div>
              <h2 className="text-2xl font-semibold text-slate-900">Reviews</h2>
              <p className="mt-2 text-sm text-slate-600">Customer feedback for this product.</p>
            </div>
          </div>
          <div className="mt-8 space-y-6">
            {reviews.length === 0 ? (
              <div className="rounded-3xl bg-slate-50 p-6 text-slate-600">No reviews yet.</div>
            ) : (
              reviews.map((review) => (
                <div key={review.id} className="rounded-3xl border border-slate-200 bg-slate-50 p-6">
                  <div className="flex items-center justify-between gap-4">
                    <div>
                      <h3 className="text-lg font-semibold text-slate-900">{review.title || 'Customer Review'}</h3>
                      <p className="mt-2 text-sm text-slate-600">{review.comment}</p>
                    </div>
                    <span className="rounded-full bg-slate-900 px-3 py-1 text-sm font-semibold text-white">{review.rating} ★</span>
                  </div>
                  {review.user_name && <p className="mt-3 text-sm text-slate-500">By {review.user_name}</p>}
                </div>
              ))
            )}
          </div>
        </div>

        <div className="rounded-[2rem] bg-white p-8 shadow-soft">
          <h2 className="text-2xl font-semibold text-slate-900">Leave a review</h2>
          {token ? (
            <form className="mt-6 space-y-6" onSubmit={async (event) => {
              event.preventDefault();
              if (!product) return;
              if (!reviewComment.trim()) {
                setReviewError('Please write a review comment.');
                return;
              }
              try {
                await addReview(product.id, {
                  rating: reviewRating,
                  title: reviewTitle,
                  comment: reviewComment
                });
                setReviews((prev) => [{
                  id: Date.now(),
                  title: reviewTitle,
                  comment: reviewComment,
                  rating: reviewRating,
                  user_name: 'You'
                }, ...prev]);
                setReviewTitle('');
                setReviewComment('');
                setReviewRating(5);
                setReviewError('');
              } catch (error) {
                console.error(error);
                setReviewError('Unable to submit review. Please try again.');
              }
            }}>
              <div>
                <label className="block text-sm font-medium text-slate-700">Rating</label>
                <select
                  value={reviewRating}
                  onChange={(e) => setReviewRating(Number(e.target.value))}
                  className="mt-3 w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
                >
                  {[5, 4, 3, 2, 1].map((value) => (
                    <option key={value} value={value}>{value} stars</option>
                  ))}
                </select>
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700">Title</label>
                <input
                  value={reviewTitle}
                  onChange={(e) => setReviewTitle(e.target.value)}
                  placeholder="Great product!"
                  className="mt-3 w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
                />
              </div>
              <div>
                <label className="block text-sm font-medium text-slate-700">Comment</label>
                <textarea
                  required
                  rows={4}
                  value={reviewComment}
                  onChange={(e) => setReviewComment(e.target.value)}
                  className="mt-3 w-full rounded-3xl border border-slate-200 bg-slate-50 px-4 py-3 text-sm outline-none focus:border-slate-900"
                />
              </div>
              {reviewError && <div className="rounded-3xl bg-rose-50 px-4 py-3 text-sm text-rose-700">{reviewError}</div>}
              <button
                type="submit"
                className="rounded-3xl bg-slate-900 px-6 py-3 text-sm font-semibold uppercase tracking-[.08em] text-white transition hover:bg-slate-700"
              >
                Submit review
              </button>
            </form>
          ) : (
            <div className="rounded-3xl bg-slate-50 p-6 text-slate-600">
              Please <a href="/login" className="font-semibold text-slate-900 underline">login</a> to leave a review.
            </div>
          )}
        </div>
      </div>
    </section>
  );
}
