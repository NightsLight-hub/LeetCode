//
// Created by sunxy on 2024/2/19.
//

#ifndef CLION_DEMO_DMI_H
#define CLION_DEMO_DMI_H
#include "atomic"
#include "dmi_base.h"
#include "thread"

using namespace std;

class SpeakerRanking {
public:
  SpeakerRanking(bool isDominant, int energyRanking, int energyScore)
      : isDominant(isDominant), energyRanking(energyRanking),
        energyScore(energyScore) {}

  bool isDominant;
  int energyRanking;
  int energyScore;

  string toString() {
    ostringstream oss;
    oss << "isDominant " << isDominant << "  energyRanking " << energyRanking
        << "  energyScore " << energyScore;
    return oss.str();
  }
};

class ActiveSpeakerChangedListener {
public:
  virtual void activeSpeakerChanged(string id, int energyScore) = 0;
};

class ActiveSpeakerDetector {
public:
  virtual void
  addActiveSpeakerChangedListener(ActiveSpeakerChangedListener *listener) = 0;

  virtual SpeakerRanking *levelChanged(string id, int level) = 0;

  virtual void removeActiveSpeakerChangedListener(
      ActiveSpeakerChangedListener *listener) = 0;
};

class AbstractActiveSpeakerDetector : ActiveSpeakerDetector {
private:
  std::recursive_mutex listenerMtx;
  vector<ActiveSpeakerChangedListener *> listeners;

public:
  void addActiveSpeakerChangedListener(
      ActiveSpeakerChangedListener *listener) override {
    if (listener == nullptr) {
      throw DmiException("listener");
    }
    std::unique_lock<std::recursive_mutex> locker(listenerMtx);
    if (find(listeners.begin(), listeners.end(), listener) == listeners.end()) {
      listeners.push_back(listener);
    }
  };

  void removeActiveSpeakerChangedListener(
      ActiveSpeakerChangedListener *listener) override {
    if (listener != nullptr) {
      std::unique_lock<std::recursive_mutex> locker(listenerMtx);
      auto it = listeners.begin();
      while (it != listeners.end()) {
        if (*it == listener) {
          listeners.erase(it);
          break;
        }
        it++;
      }
    }
  }

  void outputListeners() {
    ostringstream oss;
    oss << "[ ";
    for (ActiveSpeakerChangedListener *iter : listeners) {
      oss << "listener address " << iter << ",";
    }
    oss << " ]";
    cout << oss.str() << endl;
  }

  /**
   * 触发语音激励事件
   * @param id
   * @param energyScore
   */
  void fireActiveSpeakerChanged(const string &id, int energyScore) {
    for (ActiveSpeakerChangedListener *listener : listeners) {
      listener->activeSpeakerChanged(id, energyScore);
    }
  }
};

class DominantSpeakerIdentification : public AbstractActiveSpeakerDetector {
private:
  /**
   * The threshold of the relevant speech activities in the immediate
   * time-interval in &quot;global decision&quot;/&quot;Dominant speaker
   * selection&quot; phase of the algorithm.
   */
  constexpr static const double C1 = 3;

  /**
   * The threshold of the relevant speech activities in the medium
   * time-interval in &quot;global decision&quot;/&quot;Dominant speaker
   * selection&quot; phase of the algorithm.
   */
  constexpr static const double C2 = 2;

  /**
   * The threshold of the relevant speech activities in the long
   * time-interval in &quot;global decision&quot;/&quot;Dominant speaker
   * selection&quot; phase of the algorithm.
   */
  constexpr static const double C3 = 0;

  /**
   * The indicator which determines whether the
   * <tt>DominantSpeakerIdentification</tt> class and its instances are to
   * execute in debug mode.
   */
  constexpr static const bool DEBUG = true;

  /**
   * The interval in milliseconds of the activation of the identification of
   * the dominant speaker in a multipoint conference.
   */
  constexpr static const long DECISION_INTERVAL = 300;

  /**
   * The interval of time in milliseconds of idle execution of
   * <tt>DecisionMaker</tt> after which the latter should cease to exist. The
   * interval does not have to be very long because the background threads
   * running the <tt>DecisionMaker</tt>s are pooled anyway.
   */
  constexpr static const long DECISION_MAKER_IDLE_TIMEOUT = 15 * 1000;

  /**
   * The interval of time without a call to {@link
   * Speaker#levelChanged(int,long)} after which
   * <tt>DominantSpeakerIdentification</tt> assumes that there will be no report
   * of a <tt>Speaker</tt>'s level within a certain time-frame. The default
   * value of <tt>40</tt> is chosen in order to allow non-aggressive fading of
   * the last received or measured level and to be greater than the most common
   * RTP packet durations in milliseconds i.e. <tt>20</tt> and <tt>30</tt>.
   */
  constexpr static const long LEVEL_IDLE_TIMEOUT = 40;

  /**
   * The maximum value of audio level supported by
   * <tt>DominantSpeakerIdentification</tt>.
   */
  constexpr static const int MAX_LEVEL = 127;

  /**
   * The minimum value of audio level supported by
   * <tt>DominantSpeakerIdentification</tt>.
   */
  constexpr static const int MIN_LEVEL = 0;

  /**
   * The minimum value of speech activity score supported by
   * <tt>DominantSpeakerIdentification</tt>. The value must be positive
   * because (1) we are going to use it as the argument of a logarithmic
   * function and the latter is undefined for negative arguments and (2) we
   * will be dividing by the speech activity score.
   */
  constexpr static const double MIN_SPEECH_ACTIVITY_SCORE = 0.0000000001;

  /**
   * The threshold in terms of active sub-bands in a frame which is used
   * during the speech activity evaluation step for the medium length
   * time-interval.
   */
  constexpr static const int MEDIUM_THRESHOLD = 7;

  /**
   * The (total) number of sub-bands in the frequency range evaluated for
   * immediate speech activity. The implementation of the class
   * <tt>DominantSpeakerIdentification</tt> does not really operate on the
   * representation of the signal in the frequency domain, it works with audio
   * levels derived from RFC 6465 &quot;A Real-time Transport Protocol (RTP)
   * Header Extension for Mixer-to-Client Audio Level Indication&quot;.
   */
  constexpr static const int N1 = 13;

  /**
   * The length/size of a sub-band in the frequency range evaluated for
   * immediate speech activity. In the context of the implementation of the
   * class <tt>DominantSpeakerIdentification</tt>, it specifies the
   * length/size of a sub-unit of the audio level range defined by RFC 6465.
   */
  constexpr static const int N1_SUBUNIT_LENGTH =
      (MAX_LEVEL - MIN_LEVEL + N1 - 1) / N1;

  /**
   * The interval of time without a call to {@link
   * Speaker#levelChanged(int,long)} after which
   * <tt>DominantSpeakerIdentification</tt> assumes that a non-dominant
   * <tt>Speaker</tt> is to be automatically removed from
   * {@link #speakers}.
   */
  constexpr static const long SPEAKER_IDLE_TIMEOUT = 60 * 60 * 1000;

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
  long binomialCoefficient(int n, int r) {
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

  bool computeBigs(const Byte littles[], int littleLength, Byte bigs[],
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

  double computeSpeechActivityScore(int vL, int nR, double lambda) {
    double p = 0.5;
    double speechActivityScore = log(binomialCoefficient(nR, vL)) +
                                 vL * log(p) + (nR - vL) * log(1 - p) -
                                 log(lambda) + lambda * vL;

    if (speechActivityScore < MIN_SPEECH_ACTIVITY_SCORE) {
      speechActivityScore = MIN_SPEECH_ACTIVITY_SCORE;
    }
    return speechActivityScore;
  }

  // 用这个变量通知decision线程退出
  bool stopDecisionMaker = false;
  atomic_bool decisionMakerStart;
  thread *decisionThread{};

  /**
   * The identifier of the dominant speaker.
   */
  string dominantId = "";

  /**
   * The last/latest time at which this <tt>DominantSpeakerIdentification</tt>
   * made a (global) decision about speaker switches. The (global) decision
   * about switcher switches should be made every {@link #DECISION_INTERVAL}
   * milliseconds.
   */
  long lastDecisionTime{};

  /**
   * The time in milliseconds of the most recent (audio) level report or
   * measurement (regardless of the <tt>Speaker</tt>).
   */
  long lastLevelChangedTime{};

  /**
   * The last/latest time at which this <tt>DominantSpeakerIdentification</tt>
   * notified the <tt>Speaker</tt>s who have not received or measured audio
   * levels for a certain time (i.e. {@link #LEVEL_IDLE_TIMEOUT}) that they
   * will very likely not have a level within a certain time-frame of the
   * algorithm.
   */
  long lastLevelIdleTime{};

  /**
   * The relative speech activities for the immediate, medium and long
   * time-intervals, respectively, which were last calculated for a
   * <tt>Speaker</tt>. Simply reduces the number of allocations and the
   * penalizing effects of the garbage collector.
   */
  static const int relativeSpeechActivityLength = 3;
  double relativeSpeechActivities[relativeSpeechActivityLength]{};

  /**
   * The <tt>Speaker</tt>s in the multipoint conference associated with this
   * <tt>ActiveSpeakerDetector</tt>.
   */
  map<string, Speaker *> speakers;

  /**
   * The special ID to use to indicate silence (no speakers).
   *
   * When set to null silence detection is disabled.
   */
  const bool enableSilence;

  /**
   * The interval (in milliseconds) of no activity before switching to {@code
   * silence} (if silence detection is enabled).
   */
  const long timeoutToSilenceInterval;

  /**
   * The <tt>Speaker</tt>s in the multipoint conference with the highest
   * current energy levels.
   */
  vector<Speaker *> loudest;

  /**
   * The number of current loudest speakers to keep track of.
   */
  int numLoudestToTrack = 0;

  /**
   * Time in milliseconds after which speaker is removed from loudest list if
   * no new audio packets have been received from that speaker.
   */
  int energyExpireTimeMs = 150;

  /**
   * Alpha factor for exponential smoothing of energy values, multiplied by 100.
   */
  int energyAlphaPct = 50;

  recursive_mutex objMtx;

  /**
   * Notifies this <tt>DominantSpeakerIdentification</tt> instance that a
   * specific <tt>DecisionMaker</tt> has permanently stopped executing (in its
   * background/daemon <tt>Thread</tt>). If the specified
   * <tt>decisionMaker</tt> is the one utilized by this
   * <tt>DominantSpeakerIdentification</tt> instance, the latter will update
   * its state to reflect that the former has exited.
   *
   * @param decisionMaker the <tt>DecisionMaker</tt> which has exited
   */
  void decisionMakerExited() {
    std::unique_lock<std::recursive_mutex> locker(objMtx);
    this->decisionMakerStart.store(false);
  }

  /**
   * Retrieves a JSON representation of this instance for the purposes of the
   * REST API of Videobridge.
   * <p>
   * By the way, the method name reflects the fact that the method handles an
   * HTTP GET request.
   * </p>
   *
   * @return a <tt>JSONObject</tt> which represents this instance of the
   * purposes of the REST API of Videobridge
   */
  //      public JSONObject doGetJSON()
  //  {
  //    JSONObject jsonObject;
  //
  //    if (DEBUG)
  //    {
  //      synchronized (this)
  //      {
  //        jsonObject = new JSONObject();
  //
  //        // dominantSpeaker
  //        T dominantSpeaker = getDominantSpeaker();
  //
  //        jsonObject.put("dominantSpeaker",
  //        Objects.toString(dominantSpeaker));
  //
  //        // speakers
  //        Collection<Speaker<T>> speakersCollection = this.speakers.values();
  //        JSONArray speakersArray = new JSONArray();
  //
  //        for (Speaker<T> speaker : speakersCollection)
  //        {
  //          JSONObject speakerJSONObject = new JSONObject();
  //
  //          // id
  //          speakerJSONObject.put("id", speaker.id.toString());
  //          // levels
  //          speakerJSONObject.put("levels", speaker.getLevels());
  //          speakersArray.add(speakerJSONObject);
  //        }
  //        jsonObject.put("speakers", speakersArray);
  //      }
  //    }
  //    else
  //    {
  //      // Retrieving a JSON representation of a
  //      // DominantSpeakerIdentification has been implemented for the
  //      // purposes of debugging only.
  //      jsonObject = null;
  //    }
  //    return jsonObject;
  //  }

  /**
   * Gets the identifier of the dominant speaker.
   */
  string getDominantSpeaker() { return dominantId; }

  /**
   * Gets the <tt>Speaker</tt> in this multipoint conference identified by a
   * specific {@code id}. If no such <tt>Speaker</tt> exists, a new
   * <tt>Speaker</tt> is initialized and returned.
   *
   * @param id the identifier of the <tt>Speaker</tt> to return.
   * @return the <tt>Speaker</tt> in this multipoint conference identified by
   * {@code id}.
   */
  Speaker *getOrCreateSpeaker(const string &id) {
    std::unique_lock<std::recursive_mutex> locker(objMtx);

    Speaker *speaker = speakers[id];

    if (speaker == nullptr) {
      speaker = new Speaker(id);
      speakers[id] = speaker;

      // Since we've created a new Speaker in the multipoint conference,
      // we'll very likely need to make a decision whether there have been
      // speaker switch events soon.
      maybeStartDecisionMaker();
    }
    return speaker;
  }

  /**
   * Update loudest speaker list.
   * @param speaker the speaker with a new energy level
   * @param level the energy level
   * @param now the current time
   * @return The current ranking statistics.
   *
   * line 522
   */
  SpeakerRanking *updateLoudestList(Speaker *speaker, int level, long now) {
    bool isDominant = !dominantId.empty() && dominantId == speaker->id;

    if (level < 0) {
      /* Ignore this level, it is too old. Just gather the stats. */
      int rank = 0;
      while (rank < (int)loudest.size() && loudest[rank] != speaker)
        ++rank;

      // todo need release
      return new SpeakerRanking(isDominant, rank, speaker->energyScore);
    }

    /* exponential smoothing. round to nearest. */
    speaker->energyScore =
        (energyAlphaPct * level +
         (100 - energyAlphaPct) * speaker->energyScore + 50) /
        100;

    if (numLoudestToTrack == 0)
      return new SpeakerRanking(isDominant, 0, speaker->energyScore);

    long oldestValid = now - energyExpireTimeMs;

    cout << "Want to add " << speaker->id << " with score "
         << speaker->energyScore << ". Last level = " << level << ".";

    int i = 0;
    while (i < (int)loudest.size()) {
      Speaker *cur = loudest[i];
      if (cur->getLastLevelChangedTime() < oldestValid) {
        cout << "Removing " << cur->id << ". old.";
        loudest.erase(loudest.begin() + i);
        continue;
      }
      if (cur == speaker) {
        cout << "Removing " << cur->id << ". same.";
        loudest.erase(loudest.begin() + i);
        continue;
      }
      ++i;
    }

    int rank = 0;
    while (rank < (int)loudest.size()) {
      Speaker *cur = loudest[rank];
      if (cur->energyScore < speaker->energyScore)
        break;
      ++rank;
    }

    if (rank < numLoudestToTrack) {
      int pos = rank;
      cout << "Adding " << speaker->id << " at position " << pos << ".";
      loudest.emplace(loudest.begin() + rank, speaker);

      if ((int)loudest.size() > numLoudestToTrack)
        loudest.erase(loudest.begin() + numLoudestToTrack);
    }

    if (DEBUG) {
      i = 0;
      while (i < (int)loudest.size()) {
        Speaker *cur = loudest[i];
        cout << "New list: " << i << ": " << cur->id << ": " << cur->energyScore
             << ".";
        ++i;
      }
    }

    return new SpeakerRanking(isDominant, rank, speaker->energyScore);
  }

  /**
   * Query whether a particular endpoint is currently one of the loudest
   * speakers.
   */
public:
  /**
   * Initializes a new <tt>DominantSpeakerIdentification</tt> instance.
   * line 391
   */
  explicit DominantSpeakerIdentification(long silenceTimeout)
      : decisionMakerStart(false), enableSilence(silenceTimeout > 0),
        timeoutToSilenceInterval(silenceTimeout){};

  /**
   * Set energy ranking options
   */
  void setLoudestConfig(int numLoudestToTrack_, int energyExpireTimeMs_,
                        int energyAlphaPct_) {
    numLoudestToTrack = numLoudestToTrack_;
    energyExpireTimeMs = energyExpireTimeMs_;
    energyAlphaPct = energyAlphaPct_;
    cout << "numLoudestToTrack = " << numLoudestToTrack;
    cout << "energyExpireTimeMs = " << energyExpireTimeMs;
    cout << "energyAlphaPct = " << energyAlphaPct;

    std::unique_lock<std::recursive_mutex> locker(objMtx);
    while ((int)loudest.size() > numLoudestToTrack) {
      loudest.pop_back();
    }
    //    while (loudest.size() > numLoudestToTrack)
    //      loudest.remove(numLoudestToTrack);
  }

  bool isAmongLoudest(const string &id) {
    std::unique_lock<std::recursive_mutex> locker(objMtx);
    return any_of(loudest.begin(), loudest.end(),
                  [id](Speaker *speaker) -> bool { return speaker->id == id; });
  }

  /**
   * {@inheritDoc}
   * line 611
   */
  SpeakerRanking *levelChanged(string id, int level) override {
    Speaker *speaker;
    long now = now_milliseconds();

    {
      std::unique_lock<std::recursive_mutex> locker(objMtx);
      speaker = getOrCreateSpeaker(id);

      // Note that this ActiveSpeakerDetector is still in use. When it is
      // not in use long enough, its DecisionMaker i.e. background thread
      // will prepare itself and, consequently, this
      // DominantSpeakerIdentification for garbage collection.
      if (lastLevelChangedTime < now) {
        lastLevelChangedTime = now;

        // A report or measurement of an audio level indicates that this
        // DominantSpeakerIdentification is in use and, consequently,
        // that it'll very likely need to make a decision whether there
        // have been speaker switch events soon.
        maybeStartDecisionMaker();
      }
    }

    int cookedLevel = speaker->levelChanged(level, now);
    return updateLoudestList(speaker, cookedLevel, now);
  }

  Speaker *getSpeaker(string id) { return speakers[id]; }

  /**
   * Makes the decision whether there has been a speaker switch event. If
   * there has been such an event, notifies the registered listeners that a
   * new speaker is dominating the multipoint conference.
   */
private:
  void makeDecision(long now) {
    // If we have to fire events to any registered listeners eventually, we
    // will want to do it outside the synchronized block.
    string oldDominantSpeakerValue = "", newDominantSpeakerValue = "";
    int newDominantSpeakerEnergy;

    {
      std::unique_lock<std::recursive_mutex> locker(objMtx);
      int speakerCount = (int)speakers.size();
      string newDominantId;

      if (speakerCount == 0) {
        // If there are no Speakers in a multipoint conference, then there are
        // no speaker switch events to detect. We either have no dominant
        // speaker, or we're in a silence period (if silence detection is
        // enabled).
        newDominantId = "";
      } else if (speakerCount == 1) {
        // If there is a single Speaker in a multipoint conference, then it is
        // either the dominant speaker, or we're in a silence period (if silence
        // detection is enabled).
        Speaker *speaker = speakers.begin()->second;
        newDominantId = speaker->id;

        if (enableSilence) {
          speaker->evaluateSpeechActivityScores(now);

          long timeSinceNonSilence = now - speaker->lastNonSilence;
          if (timeSinceNonSilence > timeoutToSilenceInterval) {
            newDominantId = "";
          }
        }
      } else {
        //        bool inSilence = enableSilence && dominantId.empty();
        Speaker *dominantSpeaker =
            (dominantId.empty()) ? nullptr : speakers[dominantId];

        // If there is no dominant speaker, nominate one at random and then
        // let the other speakers compete with the nominated one.
        if (dominantSpeaker == nullptr) {
          dominantSpeaker = speakers.begin()->second;
          newDominantId = speakers.begin()->first;
        } else {
          newDominantId = dominantSpeaker->id;
        }
        // At this point dominantSpeaker==null iff inSilence==true.

        if (dominantSpeaker != nullptr) {
          dominantSpeaker->evaluateSpeechActivityScores(now);
        }

        // If multiple speakers cause speaker switches, they compete among
        // themselves by their relative speech activities in the middle
        // time-interval.
        double newDominantC2 = C2;

        for (const auto &s : speakers) {
          Speaker *speaker = s.second;

          // The dominant speaker does not compete with itself. In other
          // words, there is no use detecting a speaker switch from the
          // dominant speaker to the dominant speaker. Technically, the
          // relative speech activities are all zeroes for the dominant
          // speaker.
          if (speaker == dominantSpeaker) {
            continue;
          }

          speaker->evaluateSpeechActivityScores(now);

          // Compute the relative speech activities for the immediate,
          // medium and long time-intervals.
          for (int interval = 0; interval < relativeSpeechActivityLength;
               ++interval) {
            // When in a silence period we use MIN_SPEECH_ACTIVITY_SCORE as the
            // scores to compete against.
            double dominantSpeakerScore =
                dominantSpeaker == nullptr
                    ? MIN_SPEECH_ACTIVITY_SCORE
                    : dominantSpeaker->getSpeechActivityScore(interval);
            relativeSpeechActivities[interval] =
                log(speaker->getSpeechActivityScore(interval) /
                    dominantSpeakerScore);
          }

          double c1 = relativeSpeechActivities[0];
          double c2 = relativeSpeechActivities[1];
          double c3 = relativeSpeechActivities[2];

          if ((c1 > C1) && (c2 > C2) && (c3 > C3) && (c2 > newDominantC2)) {
            // If multiple speakers cause speaker switches, they compete
            // among themselves by their relative speech activities in
            // the middle time-interval.
            newDominantC2 = c2;
            newDominantId = s.first;
          }
        }

        if (enableSilence && dominantSpeaker != nullptr &&
            newDominantId == dominantSpeaker->id) {
          // We're not in a silence period, and none of the non-dominant
          // speakers won the challenge. Check if the current dominant speaker
          // has been silent for the timeout period, and if so switch to
          // "silence" mode.
          long timeSinceNonSilence = now - dominantSpeaker->lastNonSilence;
          if (timeSinceNonSilence > timeoutToSilenceInterval) {
            newDominantId = "";
          }
        }
      }

      if ((!newDominantId.empty() || enableSilence) &&
          newDominantId != dominantId) {
        oldDominantSpeakerValue = dominantId;
        dominantId = newDominantId;
        newDominantSpeakerValue = dominantId;
        newDominantSpeakerEnergy = speakers[dominantId]->energyScore;
      }
    } // synchronized (this)

    // Now that we are outside the synchronized block, fire events, if any,
    // to any registered listeners.
    if ((!newDominantSpeakerValue.empty() || enableSilence) &&
        newDominantSpeakerValue != oldDominantSpeakerValue) {
      fireActiveSpeakerChanged(newDominantSpeakerValue,
                               newDominantSpeakerEnergy);
    }
  }

  /**
   * Runs in the background/daemon <tt>Thread</tt> of a specific
   * <tt>DecisionMaker</tt> and makes the decision whether there has been a
   * speaker switch event.
   *
   * @param decisionMaker the <tt>DecisionMaker</tt> invoking the method
   * @return a negative integer if the <tt>decisionMaker</tt> is to exit or
   * a non-negative integer to specify the time in milliseconds until the next
   * execution of the <tt>decisionMaker</tt>
   */
  long runInDecisionMaker1() {
    {
      std::unique_lock<std::recursive_mutex> locker(objMtx);
      // If the decisionMaker has been unnecessarily executing long
      // enough, kill it in order to have a more deterministic behavior
      // with respect to disposal.
      if (0 < lastDecisionTime) {
        long idle = lastDecisionTime - lastLevelChangedTime;

        if (idle >= DECISION_MAKER_IDLE_TIMEOUT) {
          return -1;
        }
      }
    }

    return runInDecisionMaker2();
  }

  /**
   * Runs in the background/daemon <tt>Thread</tt> of {@link #decisionMaker}
   * and makes the decision whether there has been a speaker switch event.
   *
   * @return a negative integer if the <tt>DecisionMaker</tt> is to exit or
   * a non-negative integer to specify the time in milliseconds until the next
   * execution of the <tt>DecisionMaker</tt>
   */
  long runInDecisionMaker2() {
    long now = now_milliseconds();
    long levelIdleTimeout = LEVEL_IDLE_TIMEOUT - (now - lastLevelIdleTime);
    long sleep = 0;

    if (levelIdleTimeout <= 0) {
      if (lastLevelIdleTime != 0) {
        timeoutIdleLevels(now);
      }
      lastLevelIdleTime = now;
    } else {
      sleep = levelIdleTimeout;
    }

    long decisionTimeout = DECISION_INTERVAL - (now - lastDecisionTime);

    if (decisionTimeout <= 0) {
      // The identification of the dominant active speaker may be a
      // time-consuming ordeal so the time of the last decision is the
      // time of the beginning of a decision iteration.
      lastDecisionTime = now;
      makeDecision(now);
      // The identification of the dominant active speaker may be a
      // time-consuming ordeal so the timeout to the next decision
      // iteration should be computed after the end of the decision
      // iteration.
      decisionTimeout = DECISION_INTERVAL - (now_milliseconds() - now);
    }
    if ((decisionTimeout > 0) && (sleep > decisionTimeout)) {
      sleep = decisionTimeout;
    }

    return sleep;
  }

  void runDecisionMaker() {
    bool exit = false;

    while (!this->stopDecisionMaker) {
      long sleep;
      try {
        sleep = this->runInDecisionMaker1();
      } catch (exception &e) {
        // If an exception occurs we do not re-schedule.
        sleep = -1;
      }

      // A negative sleep value is contracted to mean that this DecisionMaker
      // should not re-schedule itself.
      if (sleep < 0) {
        exit = true;
      } else {
        this_thread::sleep_for(chrono::milliseconds(sleep));
      }

      if (exit) {
        this->decisionMakerExited();
        break;
      }
    }
    cout << "decisionMaker thread exit" << endl;
  }

  /**
   * Starts a background thread which is to repeatedly make the (global)
   * decision about speaker switches if such a background thread has not been
   * started yet and if the current state of this
   * <tt>DominantSpeakerIdentification</tt> justifies the start of such a
   * background thread (e.g. there is at least one <tt>Speaker</tt> in this
   * multipoint conference).
   *
   * line 817
   */
  void maybeStartDecisionMaker() {
    std::unique_lock<std::recursive_mutex> locker(objMtx);
    if ((!this->decisionMakerStart.load()) && !speakers.empty()) {
      cout << "start decisionMaker" << endl;
      decisionThread = new thread([this] { this->runDecisionMaker(); });
      decisionMakerStart.store(true);
    }
  }

  /**
   * Notifies the <tt>Speaker</tt>s in this multipoint conference who have not
   * received or measured audio levels for a certain time (i.e.
   * {@link #LEVEL_IDLE_TIMEOUT}) that they will very likely not have a level
   * within a certain time-frame of the <tt>DominantSpeakerIdentification</tt>
   * algorithm. Additionally, removes the non-dominant <tt>Speaker</tt>s who
   * have not received or measured audio levels for far too long (i.e.
   * {@link #SPEAKER_IDLE_TIMEOUT}).
   *
   * @param now the time at which the timing out is being detected
   */
  void timeoutIdleLevels(long now) {
    auto it = speakers.begin();
    while (it != speakers.end()) {
      Speaker *speaker = it->second;
      long idle = now - speaker->getLastLevelChangedTime();

      // Remove a non-dominant Speaker if he/she has been idle for far too
      // long.
      if ((SPEAKER_IDLE_TIMEOUT < idle) &&
          ((dominantId.empty()) || (speaker->id != dominantId))) {
        it = speakers.erase(it);
      } else if (LEVEL_IDLE_TIMEOUT < idle) {
        speaker->levelTimedOut();
        it++;
      } else {
        it++;
      }
    }
  }
};

#endif // CLION_DEMO_DMI_H
