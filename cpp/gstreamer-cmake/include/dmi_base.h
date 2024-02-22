//
// Created by sunxy on 2024/2/20.
//

#ifndef CLION_DEMO_DMI_BASE_H
#define CLION_DEMO_DMI_BASE_H
#include "chrono"
#include "cstring"
#include "iostream"
#include "sstream"
#include <algorithm>
#include <cmath>
#include <map>
#include <mutex>
#include <utility>
#include <vector>

typedef signed char Byte;

using namespace std;

long now_milliseconds() {
  std::chrono::time_point<std::chrono::system_clock> lastLevelChangedTime =
      chrono::system_clock::now();
  auto ms = chrono::time_point_cast<chrono::milliseconds>(lastLevelChangedTime);
  return ms.time_since_epoch().count();
}

class DmiException : public exception {
private:
  string msg;

  const char *what() const noexcept override { return msg.c_str(); }

public:
  explicit DmiException(string err) : msg(std::move(err)) {}
};

/**
 * The (total) number of long time-intervals used for speech activity score
 * evaluation at a specific time-frame.
 */
const int LONG_COUNT = 1;

/**
 * The threshold in terms of active medium-length blocks which is used
 * during the speech activity evaluation step for the long time-interval.
 */
const int LONG_THRESHOLD = 4;

/**
 * The maximum value of audio level supported by
 * <tt>DominantSpeakerIdentification</tt>.
 */
const int MAX_LEVEL = 127;

/**
 * The minimum value of audio level supported by
 * <tt>DominantSpeakerIdentification</tt>.
 */
const int MIN_LEVEL = 0;

/**
 * The number of (audio) levels received or measured for a <tt>Speaker</tt>
 * to be monitored in order to determine that the minimum level for the
 * <tt>Speaker</tt> has increased.
 */
const int MIN_LEVEL_WINDOW_LENGTH = 15 /* seconds */ * 1000 /* milliseconds */
                                    / 20 /* milliseconds per level */;

/**
 * The minimum value of speech activity score supported by
 * <tt>DominantSpeakerIdentification</tt>. The value must be positive
 * because (1) we are going to use it as the argument of a logarithmic
 * function and the latter is undefined for negative arguments and (2) we
 * will be dividing by the speech activity score.
 */
const double MIN_SPEECH_ACTIVITY_SCORE = 0.0000000001;

/**
 * The threshold in terms of active sub-bands in a frame which is used
 * during the speech activity evaluation step for the medium length
 * time-interval.
 */
const int MEDIUM_THRESHOLD = 7;

/**
 * The (total) number of sub-bands in the frequency range evaluated for
 * immediate speech activity. The implementation of the class
 * <tt>DominantSpeakerIdentification</tt> does not really operate on the
 * representation of the signal in the frequency domain, it works with audio
 * levels derived from RFC 6465 &quot;A Real-time Transport Protocol (RTP)
 * Header Extension for Mixer-to-Client Audio Level Indication&quot;.
 */
const int N1 = 13;

/**
 * The length/size of a sub-band in the frequency range evaluated for
 * immediate speech activity. In the context of the implementation of the
 * class <tt>DominantSpeakerIdentification</tt>, it specifies the
 * length/size of a sub-unit of the audio level range defined by RFC 6465.
 */
const int N1_SUBUNIT_LENGTH = (MAX_LEVEL - MIN_LEVEL + N1 - 1) / N1;

/**
 * The number of frames (i.e. {@link Speaker#immediates} evaluated for
 * medium speech activity.
 */
const int N2 = 5;

/**
 * The number of medium-length blocks constituting a long time-interval.
 */
const int N3 = 10;

class Speaker {
public:
  long lastNonSilence = -1;

private:
  static const int immediatesLength = LONG_COUNT * N3 * N2;
  Byte immediates[immediatesLength];

  /**
   * The last {@link #clock} time in millis at which this speaker was not
   * silent. We consider it being silent if
   * {@link #longSpeechActivityScore} == {@link #MIN_SPEECH_ACTIVITY_SCORE} to
   * allow short bursts of activity to not interrupt a silence period.
   */

  /**
   * The speech activity score of this <tt>Speaker</tt> for the immediate
   * time-interval.
   */
  double immediateSpeechActivityScore = MIN_SPEECH_ACTIVITY_SCORE;

  /**
   * The time in milliseconds of the most recent invocation of
   * {@link #levelChanged(int,long)} i.e. the last time at which an actual
   * (audio) level was reported or measured for this <tt>Speaker</tt>. If
   * no level is reported or measured for this <tt>Speaker</tt> long
   * enough i.e. {@link #LEVEL_IDLE_TIMEOUT}, the associated
   * <tt>DominantSpeakerIdentification</tt> will presume that this
   * <tt>Speaker</tt> was muted for the duration of a certain frame.
   */

  long lastLevelChangedTime = now_milliseconds();

  /**
   * The (history of) audio levels received or measured for this
   * <tt>Speaker</tt>.
   */
  static const int levelsLength = LONG_COUNT * N3 * N2;
  Byte levels[levelsLength];

  static const int longsLength = LONG_COUNT;
  Byte longs[longsLength];

  /**
   * The speech activity score of this <tt>Speaker</tt> for the long
   * time-interval.
   */

  double longSpeechActivityScore = MIN_SPEECH_ACTIVITY_SCORE;

  static const int mediumsLength = LONG_COUNT * N3;
  Byte mediums[LONG_COUNT * N3];

  /**
   * The speech activity score of this <tt>Speaker</tt> for the medium
   * time-interval.
   */

  double mediumSpeechActivityScore = MIN_SPEECH_ACTIVITY_SCORE;

  /**
   * The minimum (audio) level received or measured for this
   * <tt>Speaker</tt>. Since <tt>MIN_LEVEL</tt> is specified for samples
   * generated by a muted audio source, a value equal to
   * <tt>MIN_LEVEL</tt> indicates that the minimum level for this
   * <tt>Speaker</tt> has not been determined yet.
   */

  Byte minLevel = MIN_LEVEL;

  /**
   * The (current) estimate of the minimum (audio) level received or
   * measured for this <tt>Speaker</tt>. Used to increase the value of
   * {@link #minLevel}
   */

  Byte nextMinLevel = MIN_LEVEL;

  /**
   * The number of subsequent (audio) levels received or measured for this
   * <tt>Speaker</tt> which have been monitored thus far in order to
   * estimate an up-to-date minimum (audio) level received or measured for
   * this <tt>Speaker</tt>.
   */

  int nextMinLevelWindowLength;

  /** Exponential smoothing of filtered energy values.
   *  Synchronized by parent class.
   */

  /**
   * The identifier of this <tt>Speaker</tt> which is unique within this {@link
   * DominantSpeakerIdentification}.
   */
public:
  string id;
  int energyScore;
  /**
   * Initializes a new <tt>Speaker</tt> instance with a specific identifier
   *
   * @param id the object identifying this speaker.
   * instance
   */
  explicit Speaker(string id) : id(std::move(id)) {}
  recursive_mutex objMtx;
  //  mutex objMtx;

private:
  bool computeImmediates() {
    // The minimum audio level received or measured for this Speaker is
    // the level of "silence" for this Speaker. Since the various
    // Speakers may differ in their levels of "silence", put all
    // Speakers on equal footing by replacing the individual levels of
    // "silence" with the uniform level of absolute silence.
    Byte localMinLevel = (Byte)(this->minLevel + N1_SUBUNIT_LENGTH);
    bool changed = false;

    for (int i = 0; i < Speaker::immediatesLength; ++i) {
      Byte level = this->levels[i];

      if (level < localMinLevel)
        level = MIN_LEVEL;

      Byte immediate = (Byte)(level / N1_SUBUNIT_LENGTH);

      if (this->immediates[i] != immediate) {
        this->immediates[i] = immediate;
        changed = true;
      }
    }
    return changed;
  }

  static bool computeBigs(const Byte littles[], int littleLength, Byte bigs[],
                          int bigLength, int threshold) {
    int littleLengthPerBig = littleLength / bigLength;
    bool changed = false;

    for (int b = 0, l = 0; b < bigLength; b++) {
      Byte sum = 0;

      for (int lEnd = l + littleLengthPerBig; l < lEnd; l++) {
        if (littles[l] > threshold) {
          sum++;
        }
      }
      if (bigs[b] != sum) {
        bigs[b] = sum;
        changed = true;
      }
    }
    return changed;
  }

  bool computeLongs() {
    return Speaker::computeBigs(mediums, mediumsLength, longs, longsLength,
                                LONG_THRESHOLD);
  }

  bool computeMediums() {
    return computeBigs(immediates, immediatesLength, mediums, mediumsLength,
                       MEDIUM_THRESHOLD);
  }

  /**
   * Computes the binomial coefficient indexed by <tt>n</tt> and <tt>r</tt>
   * i.e. the number of ways of picking <tt>r</tt> unordered outcomes from
   * <tt>n</tt> possibilities.
   *
   * @param n the number of possibilities to pick from
   * @param r the number unordered outcomes to pick from <tt>n</tt>
   * @return the binomial coefficient indexed by <tt>n</tt> and <tt>r</tt>
   * i.e. the number of ways of picking <tt>r</tt> unordered outcomes from
   * <tt>n</tt> possibilities
   */
  static long binomialCoefficient(int n, int r) {
    int m = n - r; // r = Math.max(r, n - r);

    if (r < m) {
      r = m;
    }

    long t = 1;

    for (int i = n, j = 1; i > r; i--, j++) {
      t = t * i / j;
    }

    return t;
  }

  static double computeSpeechActivityScore(int vL, int nR, double lambda) {
    double p = 0.5;
    double speechActivityScore = log(binomialCoefficient(nR, vL)) +
                                 vL * log(p) + (nR - vL) * log(1 - p) -
                                 log(lambda) + lambda * vL;

    if (speechActivityScore < MIN_SPEECH_ACTIVITY_SCORE) {
      speechActivityScore = MIN_SPEECH_ACTIVITY_SCORE;
    }
    return speechActivityScore;
  }

  /**
   * Computes/evaluates the speech activity score of this <tt>Speaker</tt>
   * for the immediate time-interval.
   */
  void evaluateImmediateSpeechActivityScore() {
    this->immediateSpeechActivityScore =
        computeSpeechActivityScore(this->immediates[0], N1, 0.78);
    // debug
    //    cout << id << " immediateSpeechActivityScore: " <<
    //    immediateSpeechActivityScore << endl;
  }

  /**
   * Computes/evaluates the speech activity score of this <tt>Speaker</tt>
   * for the long time-interval.
   */
  void evaluateLongSpeechActivityScore(long now) {
    this->longSpeechActivityScore =
        computeSpeechActivityScore(this->longs[0], N3, 47);
    // debug
    //    cout << id << " longSpeechActivityScore: " << longSpeechActivityScore
    //    << endl;
    if (this->longSpeechActivityScore > MIN_SPEECH_ACTIVITY_SCORE) {
      this->lastNonSilence = now;
    }
  }

  /**
   * Computes/evaluates the speech activity score of this <tt>Speaker</tt>
   * for the medium time-interval.
   * line 1200
   */
  void evaluateMediumSpeechActivityScore() {
    mediumSpeechActivityScore = computeSpeechActivityScore(mediums[0], N2, 24);
    // debug
    //    cout << id << " mediumSpeechActivityScore: " <<
    //    mediumSpeechActivityScore << endl;
  }

  /**
   * Gets the (history of) audio levels received or measured for this
   * <tt>Speaker</tt>.
   *
   * @return a <tt>String</tt> that lists the (history of) audio
   * levels received or measured for this <tt>Speaker</tt>
   * line 1263
   */
  string getLevels() {
    // The levels of Speaker are internally maintained starting with the
    // last audio level received or measured for this Speaker and ending
    // with the first audio level received or measured for this Speaker.
    // Reverse the list and print them in the order they were received.
    //    Byte[] src = this.levels;
    ostringstream oss;
    oss << "[";

    for (int s = levelsLength - 1; s >= 0; --s) {
      oss << this->levels[s];
      if (s != 0) {
        oss << ",";
      } else {
        oss << "]";
      }
    }

    return oss.str();
  }

  /**
   * Updates the minimum (audio) level received or measured for this
   * <tt>Speaker</tt> in light of the receipt of a specific level.
   *
   * @param level the audio level received or measured for this
   * <tt>Speaker</tt>
   */
  void updateMinLevel(Byte level) {
    if (level != MIN_LEVEL) {
      if ((minLevel == MIN_LEVEL) || (minLevel > level)) {
        minLevel = level;
        nextMinLevel = MIN_LEVEL;
        nextMinLevelWindowLength = 0;
      } else {
        // The specified (audio) level is greater than the minimum
        // level received or measure for this Speaker. However, the
        // minimum level may be out-of-date by now. Estimate an
        // up-to-date minimum level and, eventually, make it the
        // minimum level received or measured for this Speaker.
        if (nextMinLevel == MIN_LEVEL) {
          nextMinLevel = level;
          nextMinLevelWindowLength = 1;
        } else {
          if (nextMinLevel > level) {
            nextMinLevel = level;
          }
          nextMinLevelWindowLength++;
          if (nextMinLevelWindowLength >= MIN_LEVEL_WINDOW_LENGTH) {
            // The arithmetic mean will increase the minimum
            // level faster than the geometric mean. Since the
            // goal is to track a minimum, it sounds reasonable
            // to go with a slow increase.
            double newMinLevel = sqrt(minLevel * (double)nextMinLevel);

            // Ensure that the new minimum level is within the
            // supported range.
            if (newMinLevel < MIN_LEVEL)
              newMinLevel = MIN_LEVEL;
            else if (newMinLevel > MAX_LEVEL)
              newMinLevel = MAX_LEVEL;

            minLevel = (Byte)newMinLevel;

            nextMinLevel = MIN_LEVEL;
            nextMinLevelWindowLength = 0;
          }
        }
      }
    }
  }

public:
  /**
   * Gets the speech activity score of this <tt>Speaker</tt> for a
   * specific time-interval.
   *
   * @param interval <tt>0</tt> for the immediate time-interval,
   * <tt>1</tt> for the medium time-interval, or <tt>2</tt> for the long
   * time-interval
   * @return the speech activity score of this <tt>Speaker</tt> for the
   * time-interval specified by <tt>index</tt>
   */
  double getSpeechActivityScore(int interval) const {
    switch (interval) {
    case 0:
      return immediateSpeechActivityScore;
    case 1:
      return mediumSpeechActivityScore;
    case 2:
      return longSpeechActivityScore;
    default:
      ostringstream errMsg;
      errMsg << "interval " << interval;
      throw DmiException(errMsg.str());
    }
  }

  /**
   * Evaluates the speech activity scores of this <tt>Speaker</tt> for the
   * immediate, medium, and long time-intervals. Invoked when it is time
   * to decide whether there has been a speaker switch event.
   */
  void evaluateSpeechActivityScores(long now) {
    std::unique_lock<std::recursive_mutex> locker(objMtx);
    if (computeImmediates()) {
      evaluateImmediateSpeechActivityScore();
      if (computeMediums()) {
        evaluateMediumSpeechActivityScore();
        if (computeLongs()) {
          evaluateLongSpeechActivityScore(now);
        }
      }
    }
  }
  /**
   * Gets the time in milliseconds at which an actual (audio) level was
   * reported or measured for this <tt>Speaker</tt> last.
   *
   * @return the time in milliseconds at which an actual (audio) level
   * was reported or measured for this <tt>Speaker</tt> last
   */
  long getLastLevelChangedTime() {
    std::unique_lock<std::recursive_mutex> locker(objMtx);
    return lastLevelChangedTime;
  }

  /**
   * Notifies this <tt>Speaker</tt> that a new audio level has been
   * received or measured at a specific time.
   *
   * @param level the audio level which has been received or measured for
   * this <tt>Speaker</tt>
   * @param time the (local <tt>System</tt>) time in milliseconds at which
   * the specified <tt>level</tt> has been received or measured
   * @return the audio level that was applied, after any filtering or
   * level adjustment has taken place. If negative, the audio level
   * was ignored.
   */
  int levelChanged(int level, long time) {
    std::unique_lock<std::recursive_mutex> locker(objMtx);
    // It sounds relatively reasonable that late audio levels should
    // better be discarded.
    if (lastLevelChangedTime <= time) {
      lastLevelChangedTime = time;

      // Ensure that the specified level is within the supported
      // range.
      Byte b;

      if (level < MIN_LEVEL)
        b = MIN_LEVEL;
      else if (level > MAX_LEVEL)
        b = MAX_LEVEL;
      else
        b = (Byte)level;

      // Push the specified level into the history of audio levels
      // received or measured for this Speaker.

      //      System.arraycopy(levels, 0, levels, 1, levels.length - 1);
      memmove(levels + 1, levels, (levelsLength - 1) * sizeof(Byte));
      levels[0] = b;

      // Determine the minimum level received or measured for this
      // Speaker.
      updateMinLevel(b);

      return b >= (minLevel + N1_SUBUNIT_LENGTH) ? b : 0;
    } else {
      return -1;
    }
  }

  /**
   * Notifies this <tt>Speaker</tt> that no new audio level has been
   * received or measured for a certain time which very likely means that
   * this <tt>Speaker</tt> will not have a level within a certain
   * time-frame of a <tt>DominantSpeakerIdentification</tt> algorithm.
   */
  void levelTimedOut() {
    std::unique_lock<std::recursive_mutex> locker(objMtx);
    levelChanged(MIN_LEVEL, lastLevelChangedTime);
  }
};

#endif // CLION_DEMO_DMI_BASE_H
